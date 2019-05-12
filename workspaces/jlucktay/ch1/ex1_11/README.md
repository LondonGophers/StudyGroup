# Exercise 1.11

Try `fetchall` with longer argument lists, such as samples from the top million web sites available at `alexa.com`. How
does the program behave if a web site just doesn't respond? (Section 8.9 describes mechanisms for coping in such
cases.)

## Notes

`http.Get` times out after 30 seconds, if any of the following (HTTPS issues) don't derail things first:

- x509: certificate is valid for X, not Y
- x509: certificate has expired or is not yet valid
- x509: certificate signed by unknown authority
- TLS handshake timeout
