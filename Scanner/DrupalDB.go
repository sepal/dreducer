package Scanner

import (
	"github.com/sepal/dreducer/models"
	"database/sql"
	"fmt"
	"regexp"
)

const FIELD_PATTERN = "(field_data_.+)"

type DrupalDB struct {
	db *sql.DB
	name string
	fields map[string]*models.Field
	entities map[string]*models.Entity
}

func CreateDrupalDB(database, host, user, password string) (DrupalDB, error) {
	drupal := DrupalDB{}
	dsn := fmt.Sprintf("%s:%s@%s/%s", user, password, host, database)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return drupal, err
	}

	drupal.db = db
	drupal.entities = make(map[string]*models.Entity)
	drupal.fields = make(map[string]*models.Field)

	// Scan for all tables.
	drupal.scanTables()

	// Set the correct entity ref to all field collection fields.
	drupal.scanForFC()

	return drupal, nil
}

func (d *DrupalDB) scanTables() error {
	r, _ := regexp.Compile(FIELD_PATTERN)
	rows, err := d.db.Query("SHOW TABLES")

	if err != nil {
		return err
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return err
		}

		if r.MatchString(table) {
			d.scanField(table)
		}
	}

	return nil
}

func (d *DrupalDB) scanField(table string) error {
	rows, err := d.db.Query("SELECT entity_type, bundle FROM " + table)

	if err != nil {
		return err
	}

	field, exists := d.fields[table]

	if !exists {
		f := models.CreateField(table)
		field = &f
	}

	for rows.Next() {
		var (
			entity_type string
			bundle_name string
		)

		err = rows.Scan(&entity_type, &bundle_name)

		if err != nil {
			return err
		}

		entity, exists := d.entities[entity_type]

		if !exists {
			e := models.CreateEntity(entity_type)
			entity = &e
		}

		t, exists := entity.GetType(bundle_name)

		if !exists {
			entity.AddType(bundle_name)
			t, _ = entity.GetType(bundle_name)
		}

		field.AddEntityType(t)

		entity.AddField(field, bundle_name)

		d.entities[entity_type] = entity
	}
	d.fields[field.Name] = field
	return nil
}

func (d *DrupalDB) scanForFC() {
	e, exists := d.GetEntity("field_collection_item")
	if exists {
		for _, t := range e.Types {
			f, _ := d.fields[t.Name]
			f.SetEntityTypeRef(t)
		}
	}
}

func (d *DrupalDB) GetField(name string) (*models.Field, bool) {
	f, exists := d.fields[name]
	return f, exists
}

func (d *DrupalDB) GetEntity(name string) (*models.Entity, bool) {
	e, exists := d.entities[name]
	return e, exists
}

func (d *DrupalDB) GetEntityType(entity, bundle string) (*models.EntityType, bool) {
	e, exists := d.GetEntity(entity)

	if !exists {
		return nil, false
	}

	t, exists := e.GetType(bundle)
	return t, exists
}

func (d *DrupalDB) All() ([]*models.Entity) {
	entities := make([]*models.Entity, len(d.entities))
	i := 0
	for _, e := range d.entities {
		entities[i] = e
		i++
	}
	return entities
}