#!/bin/bash

# JSON Curl Runner Script
# Usage: ./json_curl_runner.sh [folder_path] [method] [base_url]

# Default values
FOLDER_PATH="${1:-.}"
HTTP_METHOD="${2:-GET}"
BASE_URL="${3:-https://httpbin.org}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print usage
print_usage() {
    echo -e "${BLUE}Usage: $0 [folder_path] [method] [base_url]${NC}"
    echo ""
    echo "Parameters:"
    echo "  folder_path  - Path to search for JSON files (default: current directory)"
    echo "  method       - HTTP method: GET, POST, PUT, DELETE, PATCH (default: GET)"
    echo "  base_url     - Base URL for requests (default: https://httpbin.org)"
    echo ""
    echo "Examples:"
    echo "  $0                                    # GET requests from current dir"
    echo "  $0 ./data POST https://api.example.com"
    echo "  $0 /path/to/json DELETE https://api.example.com"
}

# Function to validate HTTP method
validate_method() {
    case "${1^^}" in
        GET|POST|PUT|DELETE|PATCH)
            return 0
            ;;
        *)
            echo -e "${RED}Error: Invalid HTTP method '$1'. Use GET, POST, PUT, DELETE, or PATCH${NC}"
            return 1
            ;;
    esac
}

# Function to execute curl request
execute_curl() {
    local json_file="$1"
    local method="$2"
    local url="$3"
    local filename=$(basename "$json_file")
    
    echo -e "${YELLOW}Processing: $filename${NC}"
    echo -e "${BLUE}Method: $method${NC}"
    echo -e "${BLUE}URL: $url${NC}"
    echo "----------------------------------------"
    
    case "${method^^}" in
        GET)
            # For GET requests, we'll just show the file content and make a simple GET
            echo -e "${GREEN}JSON file content:${NC}"
            cat "$json_file"
            echo ""
            echo -e "${GREEN}Executing GET request:${NC}"
            curl -s -w "\nHTTP Status: %{http_code}\nTotal Time: %{time_total}s\n" \
                -H "Accept: application/json" \
                "$url/get"
            ;;
        POST|PUT|PATCH)
            echo -e "${GREEN}Sending JSON data:${NC}"
            cat "$json_file"
            echo ""
            echo -e "${GREEN}Executing $method request:${NC}"
            curl -s -w "\nHTTP Status: %{http_code}\nTotal Time: %{time_total}s\n" \
                -X "$method" \
                -H "Content-Type: application/json" \
                -H "Accept: application/json" \
                -d "@$json_file" \
                "$url/${method,,}"
            ;;
        DELETE)
            echo -e "${GREEN}JSON file content (for reference):${NC}"
            cat "$json_file"
            echo ""
            echo -e "${GREEN}Executing DELETE request:${NC}"
            curl -s -w "\nHTTP Status: %{http_code}\nTotal Time: %{time_total}s\n" \
                -X DELETE \
                -H "Accept: application/json" \
                "$url/delete"
            ;;
    esac
    
    echo ""
    echo "========================================"
    echo ""
}

# Function to prompt user for confirmation
confirm_execution() {
    local count=$1
    echo -e "${YELLOW}Found $count JSON file(s) in '$FOLDER_PATH'${NC}"
    echo -e "${YELLOW}Method: ${HTTP_METHOD^^}${NC}"
    echo -e "${YELLOW}Base URL: $BASE_URL${NC}"
    echo ""
    read -p "Do you want to proceed? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${RED}Operation cancelled.${NC}"
        exit 0
    fi
}

# Function to list and execute JSON files
process_json_files() {
    local json_files=()
    
    # Find all JSON files in the specified folder
    while IFS= read -r -d '' file; do
        json_files+=("$file")
    done < <(find "$FOLDER_PATH" -name "*.json" -type f -print0 2>/dev/null)
    
    if [ ${#json_files[@]} -eq 0 ]; then
        echo -e "${RED}No JSON files found in '$FOLDER_PATH'${NC}"
        exit 1
    fi
    
    # List found files
    echo -e "${GREEN}Found JSON files:${NC}"
    for file in "${json_files[@]}"; do
        echo "  - $(basename "$file")"
    done
    echo ""
    
    # Ask for confirmation
    confirm_execution ${#json_files[@]}
    
    # Process each JSON file
    for json_file in "${json_files[@]}"; do
        execute_curl "$json_file" "$HTTP_METHOD" "$BASE_URL"
        
        # Add a small delay between requests
        sleep 1
    done
    
    echo -e "${GREEN}All requests completed!${NC}"
}

# Main script execution
main() {
    # Check if help is requested
    if [[ "$1" == "-h" || "$1" == "--help" ]]; then
        print_usage
        exit 0
    fi
    
    # Validate inputs
    if [ ! -d "$FOLDER_PATH" ]; then
        echo -e "${RED}Error: Directory '$FOLDER_PATH' does not exist${NC}"
        exit 1
    fi
    
    if ! validate_method "$HTTP_METHOD"; then
        exit 1
    fi
    
    # Check if curl is installed
    if ! command -v curl &> /dev/null; then
        echo -e "${RED}Error: curl is not installed${NC}"
        exit 1
    fi
    
    echo -e "${BLUE}JSON Curl Runner${NC}"
    echo "=================="
    
    process_json_files
}

# Run main function with all arguments
main "$@"
