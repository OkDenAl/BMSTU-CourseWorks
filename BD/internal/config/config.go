package config

import (
	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/cassandrinit"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/monginit"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/pkg/postgresinit"
)

type (
	Config struct {
		Postgres  postgresinit.Config `yaml:"postgres" validate:"required"`
		Mongo     monginit.Config     `yaml:"mongo" validate:"required"`
		Cassandra cassandrinit.Config `yaml:"cassandra" validate:"required"`
		Benchmark BenchmarkConfig     `yaml:"benchmark" validate:"required"`
	}

	BenchmarkConfig struct {
		ObjectsAmount   int  `yaml:"objects_amount" validate:"required"`
		CleanDataBefore bool `yaml:"clean_data_before" validate:"required"`
		CreateData      bool `yaml:"create_data" validate:"required"`
		UpdateData      bool `yaml:"update_data" validate:"required"`
		GetData         bool `yaml:"get_data" validate:"required"`
		NeedAsync       bool `yaml:"need_async"`
	}
)

func New() (*Config, error) {
	const configPath = "./config/config.yml"

	cfg := &Config{}
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	validate := validator.New()
	if err = validate.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
