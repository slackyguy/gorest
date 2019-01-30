package firebaserealtime

import (
	"log"

	"firebase.google.com/go/db"
	"github.com/slackyguy/gorest/persistence"
)

// Repository - firebase data source
type Repository struct {
	persistence.Repository
	Client *db.Client
	//app    *firebase.App
	collectionRef string
}

// SetCollectionName sets the collection name
func (repo *Repository) SetCollectionName(collection string) {
	repo.Client = Client(repo.Context, repo.AppSettings)
	repo.collectionRef = collection
}

// Find an existing item
func (repo *Repository) Find(key string, item interface{}) {
	ref := repo.Client.NewRef(repo.collectionRef).Child(key)

	if err := ref.Get(repo.Context, item); err != nil {
		log.Fatal(err)
	}
}

// List items
func (repo *Repository) List(item interface{}) {

	ref := repo.Client.NewRef(repo.collectionRef)

	if err := ref.Get(repo.Context, item); err != nil {
		log.Fatal(err)
	}

	// Lembrar da diferença entre Get e GetOrdered:
	//q := client.NewRef("accounts").OrderByChild("balance").LimitToLast(3)
	//result, err := q.GetOrdered(ctx) // Ordenado (necessário Unmarshal individual)
	//ex... if err := r.Unmarshal(&acc); err != nil {
	// Se usar o método "Get" para obter o resultado, pois "map" (coleção em
	// que o resultado é mapeado, não é ordenada
	//... if err := q.Get(ctx, &result); err != nil {

	// Transações:
	// É possível criar uma transação da seguinte forma:
	// if err := ref.Transaction(ctx, withdraw100); err != nil {
	// onde withdraw100 é uma função com a seuinte assinature:
	// func(tn db.TransactionNode) (interface{}, error) {
	// fmt.Errorf("insufficient funds: %.2f", acc.Balance)
}

// Create an item
func (repo *Repository) Create(item interface{}) string {
	ref := repo.Client.NewRef(repo.collectionRef)
	child, err := ref.Push(repo.Context, item)
	check(err)
	return child.Key
}

// Update an item
func (repo *Repository) Update(key string, item interface{}) {

	ref := repo.Client.NewRef(repo.collectionRef)
	child := ref.Child(key)

	if err := child.Set(repo.Context, item); err != nil {
		check(err)
	}
}

// Delete an item
func (repo *Repository) Delete(key string) {
	ref := repo.Client.NewRef(repo.collectionRef)
	err := ref.Child(key).Delete(repo.Context)
	check(err)
}
