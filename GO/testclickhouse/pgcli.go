package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//var pg_cl *sqlx.DB

func main() {
	db_web := "postgres://postgres:postgres@192.168.5.30:5432/nf_web"
	pg_cl, err := sqlx.Connect("postgres", db_web)
	if err != nil {
		fmt.Println(err)
	}
	pg_cl.DB.SetMaxOpenConns(100)
	pg_cl.DB.SetConnMaxIdleTime(10)
	pg_cl.DB.SetConnMaxIdleTime(0)

	db := pg_cl

	query := `SELECT usercode,gebnr FROM "public"."oa_persn"
         	WHERE "pernr" LIKE '%2800%' AND "bukrs"
			LIKE '%NVCS%' AND "departmentid" = $1 
			LIMIT 1000 OFFSET 0`
	h := []struct {
		Usercode string `db:"usercode"`
		//gebnr
		//atext
		//ename
	}{}
	err = db.Select(&h, query, "50000025")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)

	_query := `SELECT "PERNO", "PASS" 
				 FROM "AUH_USER"  
				WHERE "PERNO" = $1 and 
				      "STAT" = 'ACT' `

	// get user info from db
	_auh := []struct {
		User string `db:"PERNO"`
		Pass string `db:"PASS"`
	}{}
	db.Select(&_auh, _query, "28005669")

	fmt.Println(_auh)
}

/*
28000818	NVCS	50000025	28000818	张 鄂军	28000818		男	50	50005739				G18		28000847		中高层									19830212		湖北			已婚					D1	Z1					20190711	20220710						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工	X	Z4	20160711	G18	13636306453
28003716	NVCS	50000025	28003716	邹 国豪	28003716		男	10	50036515				G9		28000818		正式工									20010812		湖南			单身					D2	Z4					20210427	20210630						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工		Z1	20210427	G9	18569643223
28004541	NVCS	50000025	28004541	罗 佳龙	28004541		男	10	50044709				G15		28000818		正式工									19940927		安徽								D2	Z1					20211022	20241021						0000105060	NVC-MIS	1	21	I	无加无调	办公室员工	在职员工		Z1	20211022	G15	17754013578
28004567	NVCS	50000025	28004567	邓 磊	28004567		男	10	50045033				G17		28000818		正式工									19900108		湖北								D2	Z1					20211101	20241031						0000105060	NVC-MIS	1	21	I	无加无调	办公室员工	在职员工		Z1	20211101	G17	13411897029
28004672	NVCS	50000025	28004672	徐 钦	28004672		男	10	50047148				G16		28000818		中高层									19950201		湖北								D1	Z1					20211210	20241209						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工		Z1	20211210	G16	13995946368
28004680	NVCS	50000025	28004680	高 超	28004680		男	10	50047359				G14		28000818		中高层									19960220		安徽								D2	Z1					20211213	20241212						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工		Z1	20211213	G14	13063268255
28004694	NVCS	50000025	28004694	王 冬	28004694		男	10	50047556				G15		28000818		中高层									19921219		江苏								D1	Z1					20211216	20241215						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工		Z1	20211216	G15	18261611992
28004710	NVCS	50000025	28004710	闵 红剑	28004710		男	10	50047913				G16		28000818		中高层									19881222		江苏								D1	Z1					20211221	20241220						0000105060	NVC-MIS	1	10	I	无加无调	中高层管理	在职员工		Z1	20211221	G16	18717992882
28005669	NVCS	50000025	28005669	毛 雷	28005669		男		50056176				G11		28000818		正式工									19980611		湖北								D1	Z1					20221010	20251009						0000105060	NVC-MIS	1	44	I	无加无调	工程员工	在职员工		Z1	20221010	G11	18171422347
*/
