### Embedding SPA in Go binaries

This is an example of a setup for wrapping a solidjs app in a golang binary.
Three method were used, native http, echo, and fiber.

```javascript
npm run build --prefix ./client && go build ./server/<http | echo | fiber>
```
