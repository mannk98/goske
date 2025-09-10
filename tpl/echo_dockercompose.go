package tpl

func EchoServerDockerComposeTemplate() []byte {
	return []byte(`services:
  nginxgenapi:
    build:
      context: .
      dockerfile: ./Dockerfile
    image: mannk98/.{{ .AppName }}:latest
    container_name: .{{ .AppName }}
    hostname: .{{ .AppName }}
    restart: always
    networks:
      - nginxgennetwork
    ports:
      - "8080:8080"
    environment:
       LOG_LEVEL: INFO
       TZ: Asia/Ho_Chi_Minh
       PORT: 8080
       SYSTEM_SECRET: ${NGINXGENAPI_SYSTEM_SECRET}

networks:
  nginxgennetwork:
    external: true
`)
}
