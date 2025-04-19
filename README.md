# Crawler Service

A web service for crawling and parsing content from URLs, supporting both HTML pages and PDF documents.

## Features

- HTML page parsing with title, sitename, and content extraction
- PDF document downloading and text extraction
- Content conversion from HTML to Markdown
- RESTful API endpoint for crawling
- Performance timing measurements

## API Endpoints

### POST /crawl

Crawls and parses content from a given URL.

**Request Body:**
```json
{
    "url": "https://example.com/page"
}
```

**Response:**
```json
{
    "url": "https://example.com/page",
    "title": "Page Title",
    "sitename": "Example Site",
    "content": "Markdown formatted content",
    "html": "Original HTML content",
    "time": "1.234"
}
```

## Setup

1. Clone the repository
2. Configure the port in your environment
3. Run the application

## Usage

The service runs on the configured port and accepts POST requests to `/crawl` endpoint.

Example curl request:
```bash
curl -X POST http://localhost:8080/crawl \
-H "Content-Type: application/json" \
-d '{"url": "https://example.com/page"}'
```

## Error Handling

The service returns appropriate HTTP status codes:
- 200: Successful crawl
- 204: Empty URL provided
- 500: Internal server error with error message
