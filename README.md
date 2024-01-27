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
- [ ] 1. req / res 타입 정의 swagger generate
- [ ] 2. 유닛 테스트 작성
- [ ] 3. structured handle error