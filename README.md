# Use docker to create a database

```bash
sudo docker run --name some-postgres -e POSTGRES_USER=your-user -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres

# create data on docker 
sudo docker exec -it some-postgres bash

# create data on postgres
CREATE DATABASE <your-database>;
```