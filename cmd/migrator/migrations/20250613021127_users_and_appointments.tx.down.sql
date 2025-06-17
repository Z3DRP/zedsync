DROP INDEX IF EXISTS idx_role_name;
DROP TABLE IF EXISTS roles;

DROP INDEX IF EXISTS idx_usr_uid;
DROP INDEX IF EXISTS idx_usr_usrname;
DROP INDEX IF EXISTS idx_usr_email;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_company_owner;
DROP INDEX IF EXISTS idx_company_name;
DROP TABLE IF EXISTS companies;

DROP INDEX IF EXISTS idx_services_name;
DROP INDEX IF EXISTS idx_services_company;
DROP TABLE IF EXISTS services;

DROP INDEX IF EXISTS idx_appointments_user;
DROP INDEX IF EXISTS idx_appointments_service;
DROP INDEX IF EXISTS idx_appointments_times;
DROP TABLE IF EXISTS appointments;

