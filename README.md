# go-micro

## What We Build

```
프론트엔드를 포함한 5개의 MicroService 통신 
Remeber!
microservice is supposed to be self-contained.

How to Communicate Micro Services 
- Using Rest API
- Sending & Receiving using RPC
- Sending & REceivigin using gRPC
- Initiationg and responding to events using Advanced Message Queuing Protocol (AMQP)

1. Broker
   마이크로 서비스에 대한 단일 진입 지점
2. Authentication
   Postgres
3. Logger
   MongoDB
4. Mail
   이메일 발송   
5. Listener
   RabbitMQ를 통한 메시지 소비
```

## 33
Broker service to coummunicate with the logger service

## 34
Add route and handler broker to communicate with the logger service

## 35
FrontEnd modify (LoggerService Payload)

## 36
Add basic logging to the Auth service.

## 40
Add Mailer service. By Mail hog.
ㄴ 실제 SMTP를 써서 전송하는것이 아닌 호출만으로 처리를 확인가능함.

## 69
```
ERROR 이유
protoc 컴파일러가 아래 모듈을 찾지 못하는 상황.

ERROR Logs
protoc-gen-go-grpc: program not found or is not executable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.

HOW TO FIX 
protoc-gen-go-grpc 설치 후 PATH 재설정
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 75 ~101
## Use Docker Swarm
docker swarm 실습을 위한
hub.docker.com/repositories/mshero7 도커이미지 등록

도커스웜 초기세팅
```
docker swarm init 
docker swarm join-token worker
docker swarm join-token manager
docker stack deploy -c swarm.yml {{myappName}}
```

서비스 스케일아웃
```
docker service scale go-msa_listener-service=2
```

서비스 UPDATE (version up or down)
```
docker service update --image mshero7/logger-service:1.0.0 go-msa_logger-service
docker service update --image mshero7/logger-service:1.0.1 go-msa_logger-service
```

스웜에 올라온 애플리케이션 삭제.
```
docker stack rm {{myappName}}
```

front-end 도 build 한 뒤 swarm.yml 에 작성.

nginx 대신 caddy 로 reverse-proxy 처리.
