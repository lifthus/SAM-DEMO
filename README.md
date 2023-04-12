# SAM DEMO
<img src="https://user-images.githubusercontent.com/108582413/230770258-1e70e047-d7a7-4362-aefb-d529d8654d8d.png" width="30%">

**AWS SAM** Demonstration for **serverless web service** implementation in **monorepo microservice** architecture with **multiple languages**.<br>
You can get deep insight into Serverless and Infrastructure as Code. And you can also follow the practice courses below<br>

### Practice Course Korean ver.
[AWS SAM을 활용한 서버리스 웹 서비스의 구현 (1) - SAM에서 모노레포 마이크로서비스 환경 구축하기](https://medium.com/@lifthus531/aws-sam을-활용한-서버리스-웹-서비스의-구현-1-4fa3e44c99f8)<br>
[AWS SAM을 활용한 서버리스 웹 서비스의 구현 (2) - SAM에서 확장 CloudFormation으로 IaC 설정하기](https://medium.com/@lifthus531/aws-sam을-활용한-서버리스-웹-서비스의-구현-2-5874bab4badd)<br>
[AWS SAM을 활용한 서버리스 웹 서비스의 구현 (3) - SAM에서 웹 백엔드 서비스 개발하고 배포하기](https://medium.com/@lifthus531/aws-sam을-활용한-서버리스-웹-서비스의-구현-3-3718c0af556e)<br>

### Practice Course English ver.
Coming soon . . .

## Results

https://sam-demo.cloudhus.com/hello<br>
Hello-World serverless endpoint in Go

https://sam-demo.cloudhus.com/echo/wow<br>
Echo-Path serverless endpoint in Javascript

https://sam-demo.cloudhus.com/greeting/who<br>
Answering with env var serverless endpoint in Go

https://sam-demo.cloudhus.com/greeting/hi<br>
Answering with env var serverless endpoint in Go

<hr>

### Most used SAM commands
```
sam init 
```
ㄴ initializes SAM application

```
sam local generate-event apigateway aws-proxy > events/apigw-event.json
```
ㄴ generates SAM local event

```
sam local invoke -e events/apigw-event.json
```
ㄴ invokes local lambda function with event
```
sam local start-api
```
ㄴ starts local api that imitates Lambda
```
sam deploy --guided
```
ㄴ deploys with guidance
```
sam validate
```
ㄴ validates the template


### Other commands
```
sam pipeline init --bootstrap
```
ㄴ creates pipeline
```
sam sync --stack-name {stack-name} --watch
```
ㄴ tests functions in cloud
