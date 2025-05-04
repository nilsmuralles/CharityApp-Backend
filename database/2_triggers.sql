-- TRIGGER 1: Evitar administrador activo duplicado

CREATE OR REPLACE FUNCTION evitar_admin_repetido()
RETURNS TRIGGER AS $$
BEGIN
  IF EXISTS (
    SELECT 1 FROM organizacion_administradores
    WHERE id_organizacion = NEW.id_organizacion
      AND id_administrador = NEW.id_administrador
      AND fecha_fin IS NULL
  ) THEN
    RAISE EXCEPTION 'El administrador ya está activo en esta organización.';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_admin_unico
BEFORE INSERT ON organizacion_administradores
FOR EACH ROW
EXECUTE FUNCTION evitar_admin_repetido();


-- TRIGGER 2: Otorgar reconocimiento automáticamente

CREATE OR REPLACE FUNCTION otorgar_reconocimiento()
RETURNS TRIGGER AS $$
DECLARE
  total_participaciones INT;
BEGIN
  SELECT COUNT(*) INTO total_participaciones
  FROM voluntario_voluntariado
  WHERE id_voluntario = NEW.id_voluntario;

  IF total_participaciones = 5 THEN
    INSERT INTO reconocimiento (nombre, descripcion, fecha_entrega, id_voluntario)
    VALUES (
      'Voluntario Activo',
      'Participación destacada en al menos 5 voluntariados',
      CURRENT_DATE,
      NEW.id_voluntario
    );
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_reconocimiento
AFTER INSERT ON voluntario_voluntariado
FOR EACH ROW
EXECUTE FUNCTION otorgar_reconocimiento();


-- TRIGGER 3: Cerrar voluntariado si todos terminaron

-- Asegurar que la tabla voluntariado tenga un campo de estado:
ALTER TABLE voluntariado ADD COLUMN estado VARCHAR(20) DEFAULT 'Activo';

CREATE OR REPLACE FUNCTION cerrar_voluntariado_si_completo()
RETURNS TRIGGER AS $$
DECLARE
  total INT;
  finalizados INT;
BEGIN
  SELECT COUNT(*) INTO total
  FROM voluntario_voluntariado
  WHERE id_voluntariado = NEW.id_voluntariado;

  SELECT COUNT(*) INTO finalizados
  FROM voluntario_voluntariado
  WHERE id_voluntariado = NEW.id_voluntariado AND hora_fin IS NOT NULL;

  IF total > 0 AND total = finalizados THEN
    UPDATE voluntariado
    SET estado = 'Finalizado'
    WHERE id = NEW.id_voluntariado;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_cerrar_voluntariado
AFTER UPDATE OF hora_fin ON voluntario_voluntariado
FOR EACH ROW
WHEN (NEW.hora_fin IS NOT NULL)
EXECUTE FUNCTION cerrar_voluntariado_si_completo();

