package usecase

import "github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/config"

func (b Benchmark) buildBenchFuncs(cfg config.BenchmarkConfig) []BenchFunc {
	benchFuncs := make([]BenchFunc, 0)

	if cfg.CreateData {
		if cfg.NeedAsync {
			benchFuncs = append(benchFuncs, b.createBenchAsync)
		} else {
			benchFuncs = append(benchFuncs, b.createBench)
		}
	}

	if cfg.UpdateData {
		if cfg.NeedAsync {
			benchFuncs = append(benchFuncs, b.updateBenchAsync)
		} else {
			benchFuncs = append(benchFuncs, b.updateBench)
		}
	}

	if cfg.GetData {
		if cfg.NeedAsync {
			benchFuncs = append(benchFuncs, b.getBenchAsync)
		} else {
			benchFuncs = append(benchFuncs, b.getBench)
		}
	}

	return benchFuncs
}
