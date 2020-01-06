package dataframe_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func generateSeries(n, rep int) (data []series.Series) {
	rand.Seed(100)
	for j := 0; j < rep; j++ {
		var is []int
		var bs []bool
		var fs []float64
		var ss []string
		for i := 0; i < n; i++ {
			is = append(is, rand.Int())
		}
		for i := 0; i < n; i++ {
			fs = append(fs, rand.Float64())
		}
		for i := 0; i < n; i++ {
			ss = append(ss, strconv.Itoa(rand.Int()))
		}
		for i := 0; i < n; i++ {
			r := rand.Intn(2)
			b := false
			if r == 1 {
				b = true
			}
			bs = append(bs, b)
		}
		data = append(data, series.Ints(is))
		data = append(data, series.Bools(bs))
		data = append(data, series.Floats(fs))
		data = append(data, series.Strings(ss))
	}
	return
}

func generateNamedSeries(n, rep int) (data []series.Series) {
	rand.Seed(100)
	for j := 0; j < rep; j++ {
		var is []int
		var bs []bool
		var fs []float64
		var ss []string
		for i := 0; i < n; i++ {
			is = append(is, rand.Intn(n))
		}
		for i := 0; i < n; i++ {
			fs = append(fs, rand.Float64())
		}
		for i := 0; i < n; i++ {
			ss = append(ss, strconv.Itoa(rand.Int()))
		}
		for i := 0; i < n; i++ {
			r := rand.Intn(2)
			b := false
			if r == 1 {
				b = true
			}
			bs = append(bs, b)
		}
		data = append(data, series.New(is, series.Int, "col1"))    // series.Ints(is))
		data = append(data, series.New(bs, series.Bool, "col2"))   //series.Bools(bs))
		data = append(data, series.New(fs, series.Float, "col3"))  //series.Floats(fs))
		data = append(data, series.New(ss, series.String, "col4")) //series.Strings(ss))
	}
	return
}

func generateIntsN(n, k int) (data []int) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(k))
	}
	return
}

func BenchmarkNew(b *testing.B) {
	table := []struct {
		name string
		data []series.Series
	}{
		{
			"100000x4",
			generateSeries(100000, 1),
		},
		{
			"100000x40",
			generateSeries(100000, 10),
		},
		{
			"100000x400",
			generateSeries(100000, 100),
		},
		{
			"1000x40",
			generateSeries(1000, 10),
		},
		{
			"1000x4000",
			generateSeries(1000, 1000),
		},
		{
			"1000x40000",
			generateSeries(1000, 10000),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dataframe.New(test.data...)
			}
		})
	}
}

func BenchmarkDataFrame_Arrange(b *testing.B) {
	data := dataframe.New(generateSeries(100000, 5)...)
	table := []struct {
		name string
		data dataframe.DataFrame
		key  []dataframe.Order
	}{
		{
			"100000x20_1",
			data,
			[]dataframe.Order{dataframe.Sort("X0")},
		},
		{
			"100000x20_2",
			data,
			[]dataframe.Order{
				dataframe.Sort("X0"),
				dataframe.Sort("X1"),
			},
		},
		{
			"100000x20_3",
			data,
			[]dataframe.Order{
				dataframe.Sort("X0"),
				dataframe.Sort("X1"),
				dataframe.Sort("X2"),
			},
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.data.Arrange(test.key...)
			}
		})
	}
}

func BenchmarkDataFrame_Subset(b *testing.B) {
	b.ReportAllocs()
	data1000x20 := dataframe.New(generateSeries(1000, 5)...)
	data1000x200 := dataframe.New(generateSeries(1000, 50)...)
	data1000x2000 := dataframe.New(generateSeries(1000, 500)...)
	data100000x20 := dataframe.New(generateSeries(100000, 5)...)
	data1000000x20 := dataframe.New(generateSeries(1000000, 5)...)
	idx10 := generateIntsN(10, 10)
	idx100 := generateIntsN(100, 100)
	idx1000 := generateIntsN(1000, 1000)
	idx10000 := generateIntsN(10000, 10000)
	idx100000 := generateIntsN(100000, 100000)
	idx1000000 := generateIntsN(1000000, 1000000)
	table := []struct {
		name    string
		data    dataframe.DataFrame
		indexes interface{}
	}{
		{
			"1000000x20_100",
			data1000000x20,
			idx100,
		},
		{
			"1000000x20_1000",
			data1000000x20,
			idx1000,
		},
		{
			"1000000x20_10000",
			data1000000x20,
			idx10000,
		},
		{
			"1000000x20_100000",
			data1000000x20,
			idx100000,
		},
		{
			"1000000x20_1000000",
			data1000000x20,
			idx1000000,
		},
		{
			"100000x20_100",
			data100000x20,
			idx100,
		},
		{
			"100000x20_1000",
			data100000x20,
			idx1000,
		},
		{
			"100000x20_10000",
			data100000x20,
			idx10000,
		},
		{
			"100000x20_100000",
			data100000x20,
			idx100000,
		},
		{
			"1000x20_10",
			data1000x20,
			idx10,
		},
		{
			"1000x20_100",
			data1000x20,
			idx100,
		},
		{
			"1000x20_1000",
			data1000x20,
			idx1000,
		},
		{
			"1000x200_10",
			data1000x200,
			idx10,
		},
		{
			"1000x200_100",
			data1000x200,
			idx100,
		},
		{
			"1000x200_1000",
			data1000x200,
			idx1000,
		},
		{
			"1000x2000_10",
			data1000x2000,
			idx10,
		},
		{
			"1000x2000_100",
			data1000x2000,
			idx100,
		},
		{
			"1000x2000_1000",
			data1000x2000,
			idx1000,
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.data.Subset(test.indexes)
			}
		})
	}
}

func BenchmarkDataFrame_Elem(b *testing.B) {
	data := dataframe.New(generateSeries(100000, 5)...)
	table := []struct {
		name string
		data dataframe.DataFrame
	}{
		{
			"100000x20_ALL",
			data,
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for k := 0; k < 100000; k++ {
					test.data.Elem(k, 0)
				}
			}
		})
	}
}

func BenchmarkDataFrame_InnerJoin(b *testing.B) {
	// use `benchtime` if required
	// e.g.: go test -run=X -bench=InnerJoin -benchtime=100s
	tests := []struct {
		name     string
		desc     string
		df1      dataframe.DataFrame
		df2      dataframe.DataFrame
		keys     []string
		expected string
	}{
		{
			"InnerJoin_3_X_1",
			"",
			dataframe.New(
				series.New([]string{"111", "222", "333"}, series.String, "id"),
				series.New([]string{"aaron", "bey", "ryan"}, series.String, "name"),
			),
			dataframe.New(
				series.New([]string{"222"}, series.String, "id"),
				series.New([]string{"45"}, series.String, "age"),
			),
			[]string{"id"},
			``,
		},
		{
			"InnerJoin_10_X_10",
			"",
			dataframe.New(generateNamedSeries(10, 1)...),
			dataframe.New(generateNamedSeries(10, 1)...),
			[]string{"col1"},
			``,
		},
		{
			"InnerJoin_10k_X_10k",
			"",
			dataframe.New(generateNamedSeries(10000, 1)...),
			dataframe.New(generateNamedSeries(10000, 1)...),
			[]string{"col1"},
			``,
		},
	}
	for _, test := range tests {
		var x, y int
		// df := dataframe.New(generateSeries(1, 1)...)
		// fmt.Println("df1\n", test.df1)
		// fmt.Println("df2\n", test.df2)
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				df := test.df1.InnerJoin(test.df2, test.keys...)
				x, y = df.Dims()
			}
		})
		// fmt.Println("df\n", df)
		fmt.Println("Resulting df size = ", x, "X", y)
	}
}
