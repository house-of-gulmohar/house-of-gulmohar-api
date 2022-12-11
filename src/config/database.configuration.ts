import { registerAs } from '@nestjs/config';

export default registerAs('database', () => ({
  host: process.env.DATABSE_HOST || 'localhost',
  port: process.env.DATABASE_PORT || 5432,
  db: process.env.DATABASE_DB,
  user: process.env.DATABASE_USER || 'postgres',
  password: process.env.DATABASE_PASSWORD,
}));
