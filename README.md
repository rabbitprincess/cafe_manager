# cafe manager

manager program for cafe

## Build

    git clone https://github.com/gokch/cafe_manager
    cd cafe_manager && make

## Run With flags

```
Usage:
  ./bin/cafe_manager [flags]

Flags:
  -A, --dbaddr string   db address (default "localhost")
  -n, --dbname string   db name (default "cafe")
  -c, --dbpass string   db password (default "1234")
  -P, --dbport string   db port (default "3306")
  -u, --dbuser string   db user name (default "root")
  -h, --help            help for cafe_manager
  -p, --port string     port (default "3000")
```

## Build and run using Docker Compose

    git clone https://github.com/gokch/cafe_manager
    cd cafe_manager && docker compose up

# 해야될거 ( 임시 )
- [ ] 1. jwt 커스터마이징 및 테스트
- [ ] 2. req / res 타입 정의 swagger generate
- [ ] 3. 유닛 테스트 작성
- [ ] 4. 에러 핸들링 ( 시간남으면 )
- [ ] 5. cli 작성 ( cobra 사용 )
- [ ] 5. docker compose 작성 ( go + mysql 실행 -> mysql 스키마 초기화 로직 필요 )