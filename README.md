# Debt-tracker-backend
The application, named Debt-tracker, primarily targets individuals residing in shared apartments who collectively manage household expenditures like groceries, toiletries, and more. This tool equally distributes costs among all apartment residents and computes the individualized debt allocation for each participant.

Backend technologies and frameworks [(demo)](https://www.youtube.com/watch?v=jfPyikD0FkY):
- Gin
- Makefile
- Docker & docker compose
- PostgreSQL

### How to run
```
export DB_PASSWORD={YoUr_PaSsWoRd}
docker-compose stop
docker-compose rm -f
docker-compose pull
docker-compose up -d
```
