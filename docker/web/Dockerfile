# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:latest

# Create a directory inside the container
RUN mkdir -p /go/src/github.com/WymA/go-restful
WORKDIR /go/src/github.com/WymA/go-restful

# Copy the project into container
COPY . /go/src/github.com/WymA/go-restful

# Download and install any required third party dependencies into the container.
# RUN go get github.com/codegangsta/gin
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/go-sql-driver/mysql
# RUN go get gopkg.in/gorp.v1
# RUN go-wrapper download
# RUN go-wrapper install

# Set the PORT environment variable inside the container
# ??????????????????????????????????????????????????????
# ENV PORT 8080

# Expose port 3000 to the host so we can access the gin proxy
EXPOSE 3000

# Now tell Docker what command to run when the container starts
CMD go run main.go
