package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	pb "github.com/Sandhya-Pratama/Libary-API/book-service/book-service/common/proto"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/builder"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/config"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/binder"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/server"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/http/validator"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/repository"
	"github.com/Sandhya-Pratama/Libary-API/book-service/internal/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//membuat config ke env
	cfg, err := config.NewConfig(".env")
	checkError(err)

	//mengubah starter
	splash()

	//membuat koneksi ke postgres
	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	//membuat koneksi ke redis
	// redisClient := buildRedis(cfg)

	//membuat routes
	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)

	//membuat binder
	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	//membuat server
	svr := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	//membuat start server
	runServer(svr, cfg.Port)

	// Menjalankan gRPC server pada port lain, misalnya 50051
	runGRPCServer(db)

	//membuat shutdown
	waitForShutdown(svr)

}

// func untuk start server
func runServer(srv *server.Server, port string) {
	go func() {
		// Pastikan port diawali dengan ":"
		fmt.Printf("Starting server on port: %s\n", port)
		err := srv.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}
	}()
}

// funv untuk start gRPC server
func runGRPCServer(db *gorm.DB) {
	go func() {
		// Membuat gRPC server
		bookRepo := repository.NewBookRepository(db)
		bookService := service.NewBookService(bookRepo)
		grpcServer := grpc.NewServer()
		pb.RegisterBookServiceServer(grpcServer, service.NewBookGRPCServer(bookService))

		listener, err := net.Listen("tcp", ":50051") // gRPC server berjalan di port 50051
		if err != nil {
			log.Fatalf("Failed to listen on port 50051: %v", err)
		}
		fmt.Println("Starting gRPC server on port 50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
}

// func untuk shutdown
func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

// func untuk koneksi ke postgresql
func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ini buar ngecek gormnya
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func untuk koneksi ke redis
// func buildRedis(cfg *config.Config) *redis.Client {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
// 		Password: cfg.Redis.Password,
// 		DB:       0,
// 	})
// 	return client
// }

// untuk membuat spalsh menggunakan patorjk.com
func splash() {
	colorReset := "\030[0m"

	splashText := `
.____    ._____.                                          _____                 
|    |   |__\_ |______________ _______ ___.__.           /  _  \ ______ ______  
|    |   |  || __ \_  __ \__  \\_  __ <   |  |  ______  /  /_\  \\____ \\____ \ 
|    |___|  || \_\ \  | \// __ \|  | \/\___  | /_____/ /    |    \  |_> >  |_> >
|_______ \__||___  /__|  (____  /__|   / ____|         \____|__  /   __/|   __/ 
        \/       \/           \/       \/                      \/|__|   |__|                                          |_|   |_|        
`
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

// func untuk cek error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
