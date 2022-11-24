// package main
//
// import "fmt"
//
// func main() {
//     fmt.Println("Hello world")
// }

package main

import (
  "fmt"
  "log"
  "context"
  "time"
  "cloud.google.com/go/firestore"
//   firebase "firebase.google.com/go"
//   "google.golang.org/api/option"
)

// func main() {
//
// //   conf := &firebase.Config{ProjectID: projectID}
// //   app, err := firebase.NewApp(context.Background(), conf)
//
//     // Use the application default credentials
//     ctx := context.Background()
//     conf := &firebase.Config{ProjectID: 'strongman-295718'}
//     app, err := firebase.NewApp(ctx, conf)
//     if err != nil {
//       log.Fatalln(err)
//     }
//
//     client, err := app.Firestore(ctx)
//     if err != nil {
//       log.Fatalln(err)
//     }
//     // defer client.Close()
//
//     fmt.Println("Hello world")
//
//     defer client.Close()
// }


func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "strongman-295718"

	// [END firestore_setup_client_create]
	// Override with -project flags
    // 	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
    // 	flag.Parse()

	// [START firestore_setup_client_create]
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

var start time.Time

func init() {
    start = time.Now()
}

func main() {
    fmt.Println("main execution started at time", time.Since(start))
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	dsnap, err := client.Collection("lifts").Doc("02efca3d5116829dd1069d41f88ef3f3").Get(ctx)
    //     if err != nil {
    //         return nil, err
    //     }
    if err != nil {
        log.Fatalln(err)
    }
    m := dsnap.Data()
    fmt.Printf("Document data: %#v\n", m)
    fmt.Println("\nmain execution stopped at time", time.Since(start))
    // 	// [START firestore_setup_dataset_pt1]
    // 	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
    // 		"first": "Ada",
    // 		"last":  "Lovelace",
    // 		"born":  1815,
    // 	})
    // 	if err != nil {
    // 		log.Fatalf("Failed adding alovelace: %v", err)
    // 	}
    // 	// [END firestore_setup_dataset_pt1]
    //
    // 	// [START firestore_setup_dataset_pt2]
    // 	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
    // 		"first":  "Alan",
    // 		"middle": "Mathison",
    // 		"last":   "Turing",
    // 		"born":   1912,
    // 	})
    // 	if err != nil {
    // 		log.Fatalf("Failed adding aturing: %v", err)
    // 	}
    // 	// [END firestore_setup_dataset_pt2]
    //
    // 	// [START firestore_setup_dataset_read]
    // 	iter := client.Collection("users").Documents(ctx)
    // 	for {
    // 		doc, err := iter.Next()
    // 		if err == iterator.Done {
    // 			break
    // 		}
    // 		if err != nil {
    // 			log.Fatalf("Failed to iterate: %v", err)
    // 		}
    // 		fmt.Println(doc.Data())
    // 	}
	// [END firestore_setup_dataset_read]
}
// func getQuote() *Quote {
//   "quote": "Play it",
//   "author": "Casablanca"
// }