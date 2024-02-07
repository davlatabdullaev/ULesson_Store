package postgres

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
	"test/config"
	"test/storage"

	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	_ "github.com/lib/pq"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("error while parsing config", err.Error())
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	// migration
	m, err := migrate.New("file://migrations/postgres/", url)
	if err != nil {
		fmt.Println("error while migrating", err.Error())
		return nil, err
	}

	if err = m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			if err != nil {
				fmt.Println("err in checking version and dirty", err.Error())
				return nil, err
			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					fmt.Println("ERR in making force", err.Error())
					return nil, err
				}
			}
			fmt.Println("ERROR in migrating", err.Error())
			return nil, err
		}
	}

	return Store{
		pool: pool,
	}, nil
}

func (s Store) Close() {
	s.pool.Close()
}

func (s Store) User() storage.IUserStorage {
	return NewUserRepo(s.pool)
}

func (s Store) Category() storage.ICategoryStorage {
	return NewCategoryRepo(s.pool)
}

func (s Store) Product() storage.IProductStorage {
	return NewProductRepo(s.pool)
}

func (s Store) Basket() storage.IBasketStorage {
	return NewBasketRepo(s.pool)

}

func (s Store) BasketProduct() storage.IBasketProductStorage {
	return NewBasketProductRepo(s.pool)
}

func (s Store) Store() storage.IStoreStorage {
	return NewStoreRepo(s.pool)
}

func (s Store) Branch() storage.IBranchStorage {
	return NewBranchRepo(s.pool)
}

func (s Store) Dealer() storage.IDealerStorage {
	return NewDealerRepo(s.pool)
}

func (s Store) Income() storage.IIncomeStorage {
	return NewIncomeRepo(s.pool)
}

func (s Store) IncomeProduct() storage.IIncomeProductStorage {
	return NewIncomeProductRepo(s.pool)
}