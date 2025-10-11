APP_NAME=authservice

generation:
	echo "🚀 Generating templates for authentication..."
	cd authentication && go run cmd/template/main.go
	echo "✅ Done with authentication"
	echo "🚀 Generating templates for authorization..."
	cd authorization && go run cmd/template/main.go
	echo "✅ Done with authorization"
	echo "🚀 Generating proto file..."
	sh ./generate_protos.sh
	echo "✅ Done with proto"

