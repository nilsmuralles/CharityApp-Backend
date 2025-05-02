CREATE TABLE administrador (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL CHECK (fecha_nacimiento <= CURRENT_DATE),
    correo VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE donador (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL CHECK (fecha_nacimiento <= CURRENT_DATE),
    correo VARCHAR(100) NOT NULL UNIQUE,
    categoria VARCHAR(50) NOT NULL CHECK (categoria IN ('Individual', 'Empresa', 'Anonimo'))
);

CREATE TABLE voluntario (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL CHECK (fecha_nacimiento <= CURRENT_DATE),
    correo VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE organizacion (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,
    correo VARCHAR(100) NOT NULL UNIQUE,
    direccion VARCHAR(200) NOT NULL
);

CREATE TABLE organizacion_administradores (
    id SERIAL PRIMARY KEY,
    id_organizacion INT NOT NULL REFERENCES organizacion(id),
    id_administrador INT NOT NULL REFERENCES administrador(id),
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE,
    UNIQUE (id_organizacion, id_administrador)
);

CREATE TABLE campania (
    id SERIAL PRIMARY KEY,
    id_organizacion INT NOT NULL REFERENCES organizacion(id),
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    CHECK (fecha_inicio <= fecha_fin)
);

CREATE TABLE donacion_monetaria (
    id SERIAL PRIMARY KEY,
    id_donador INT NOT NULL REFERENCES donador(id),
    id_campania INT NOT NULL REFERENCES campania(id),
    monto NUMERIC(10,2) NOT NULL CHECK (monto > 0),
    fecha DATE NOT NULL DEFAULT CURRENT_DATE,
    metodo_pago VARCHAR(50) NOT NULL CHECK (metodo_pago IN ('Tarjeta', 'Transferencia', 'Efectivo'))
);

CREATE TABLE donacion_no_monetaria (
    id SERIAL PRIMARY KEY,
    id_donador INT NOT NULL REFERENCES donador(id),
    id_campania INT NOT NULL REFERENCES campania(id),
    fecha DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE articulo_donado (
    id SERIAL PRIMARY KEY,
    id_donacion_no_monetaria INT NOT NULL REFERENCES donacion_no_monetaria(id) ON DELETE CASCADE,
    nombre VARCHAR(100) NOT NULL,
    cantidad INT NOT NULL CHECK (cantidad > 0)
);

CREATE TABLE voluntariado (
    id SERIAL PRIMARY KEY,
    fecha DATE NOT NULL DEFAULT CURRENT_DATE,
    id_campania INT NOT NULL REFERENCES campania(id)
);

CREATE TABLE evento (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    fecha DATE NOT NULL,
    lugar VARCHAR(100) NOT NULL,
    id_voluntariado INT NOT NULL REFERENCES voluntariado(id)
);

CREATE TABLE reconocimiento (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    fecha_entrega DATE NOT NULL,
    id_voluntario INT NOT NULL REFERENCES voluntario(id)
);

CREATE TABLE voluntario_voluntariado (
    id SERIAL PRIMARY KEY,
    id_voluntario INT NOT NULL REFERENCES voluntario(id),
    id_voluntariado INT NOT NULL REFERENCES voluntariado(id),
    hora_inicio TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    hora_fin TIMESTAMP,
    CHECK (hora_fin IS NULL OR hora_inicio <= hora_fin),
    UNIQUE (id_voluntario, id_voluntariado)
);
