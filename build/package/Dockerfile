FROM ubuntu:latest

#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh
RUN     apt-get  update -y \
&& apt-get upgrade -y \
&& apt-get install iputils-ping -y \
&& apt-get install net-tools -y \
&& apt-get install bash -y

# Set the Current Working Directory inside the container
WORKDIR ../..

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
CMD ["ls -la"]
# Build the Go app
#RUN go build -o /cmd/url_shortener/main .

# Expose port 8080 to the outside world
#EXPOSE 8080

# Run the executable
#CMD ["./main"]
CMD ["bash"]
#CMD bash