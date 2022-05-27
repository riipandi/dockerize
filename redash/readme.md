# Read Only PostgreSQL User

### 1. Create a read only role
```sql
CREATE ROLE myapp_readonly;
GRANT CONNECT ON DATABASE defaultdb TO myapp_readonly;
GRANT USAGE ON SCHEMA myapp TO myapp_readonly;
```

### 2. Grant permissions to the new role
```sql
GRANT SELECT ON TABLE "myapp"."employees" TO myapp_readonly;
GRANT SELECT ON TABLE "myapp"."jobs" TO myapp_readonly;
GRANT SELECT (id, name) ON TABLE myapp.customers TO myapp_readonly;
```

### 3. Create user for Redash and assign it the read only role
```sql
CREATE USER redash WITH PASSWORD 'secret';
GRANT myapp_readonly TO redash;
```
