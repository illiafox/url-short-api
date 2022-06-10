package links

import (
	"context"
	"testing"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stretchr/testify/suite"
	"ozon-url-shortener/app/internal/config"
	"ozon-url-shortener/app/internal/domain/links"
	"ozon-url-shortener/app/pkg/client/pg"
	"ozon-url-shortener/app/pkg/generator"
)

type DBSuite struct {
	suite.Suite

	storage links.Storage
	done    func(context.Context) error
}

func (suite *DBSuite) SetupSuite() {
	r := suite.Require()
	//
	cfg := config.Postgres{
		User:     "user",
		Pass:     "password",
		DBName:   "links_test",
		IP:       "127.0.0.1",
		Port:     5432,
		Protocol: "tcp",
	}
	//
	err := cleanenv.ReadEnv(&cfg)
	r.NoError(err, "read environment variables")
	//
	pool, err := pg.NewPool(5*time.Second, pg.Config(cfg))
	r.NoError(err, "new pool")
	// //
	{
		ctx := context.Background()

		var count int
		err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM links").Scan(&count)
		r.NoError(err, "select count from links")

		r.Equal(0, count, "'links' table is not empty")
	}
	// //
	suite.storage = NewPgStorage(pool)
	suite.done = func(ctx context.Context) error {
		_, err := pool.Exec(ctx, "TRUNCATE TABLE links")
		pool.Close()

		return err
	}
}

func (suite *DBSuite) TearDownSuite() {
	if suite.storage == nil {
		return
	}
	//
	ctx := context.Background()
	err := suite.done(ctx)
	suite.Require().NoError(err, "truncate tables")
}

func (suite *DBSuite) TestGetSet() {
	ctx := context.Background()
	a := suite.Assert()
	//
	url := "https://github.com/illiafox"
	key, err := generator.Key(10)
	a.NoError(err, "generate key")
	//
	err = suite.storage.StoreURL(ctx, key, url)
	a.NoError(err, "store url %s with key %s", url, key)
	//
	data, err := suite.storage.GetURL(ctx, string(key))
	a.NoError(err, "get url from key %s", key)
	a.Equal(url, data, "compare url")
}

func TestDB(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DBSuite))
}
