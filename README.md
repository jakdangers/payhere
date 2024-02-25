# 백엔드 엔지니어 과제 (과제 내용을 직접적으로 복붙하지 않음)

유저는 회원가입을 할 수 있고 로그인 후에는 상품 등록, 조회, 부분 수정, 삭제, 리스트 조회 REST API

프로젝트 실행

```bash
make compose-up       
```

프로젝트 종료 (과제 앱의 이미지만 삭제하고 mysql 5.7은 삭제 하지 않습니다)

```bash
make compose-clean     
```

## 단위 테스트

```bash
make test
```

### API 테스트 (API SPEC스팩은 스웨거)

원할한 과제 테스트를 위해 swageer의 example을 추가하여 스웨거상에서 바로 실행 할 수 있도록 했습니다.

또한 도커 컴포즈 실행시 init.sql을 통해 테스트 데이터를 삽입하였습니다.

상품 32개를 갖고 있는 사장님
```bash
로그인 정보
{
  "mobileID": "01011111111",
  "password": "1234"
}
```

다른 유저의 상품에 접근하지 못하는 것을 테스트하기 위한 유저 정보
```bash
로그인 정보
{
  "mobileID": "01022222222",
  "password": "1234"
}
```


```python
import foobar

# returns 'words'
foobar.pluralize('word')

# returns 'geese'
foobar.pluralize('goose')

# returns 'phenomenon'
foobar.singularize('phenomena')
```
# 구현 코드에 대한 생각

### 프로젝트 실행

여러 개의 도커 이미지를 연결하여 띄우려면 Docker Compose를 사용하는 것이 일반적인 방법이라고 생각했습니다. 요구사항에 도커 컴포즈를 사용하라고 명시되어 있진 않았지만 두 개의 서비스를 실행하고 서로 통신하기에 용이하다고 생각했습니다.

### 프로젝트 구조

[고랭 클린 아키텍처](https://github.com/bxcodec/go-clean-arch)을 참고해서 프로젝트 구조를 만들었습니다.

![](https://velog.velcdn.com/images/jakdangers/post/87659488-6177-4e46-a793-1834abbd3e87/image.png)

### 추가 패키지 사용

	github.com/DATA-DOG/go-sqlmock - repository 테스트
	github.com/go-sql-driver/mysql - 디비 커넥트
	github.com/golang-jwt/jwt/v5 - jwt 토큰 인증
	github.com/spf13/viper - 설정 yaml 읽기
	github.com/stretchr/testify - 단위 테스트를 위한 assert, mock 사용
	github.com/swaggo/files - api 문서
	github.com/swaggo/gin-swagger - api 문서
	github.com/swaggo/swag - api 문서
	golang.org/x/crypto v0.19.0 - 비밀번호 암호화
	k8s.io/utils v0.0.0-20240102154912-e7106e64919e - 프리미티드 타입 포인터로 변환

gin 외에 과제를 구현하기 위해 최소한의 외부 의존성을 사용하려고 했습니다. 라이브러리를 최소한으로 사용하고 코드적으로 풀어나가는게 방향성에 맞을 것 같아
필수적인 것을 제외하고 설정을 읽는 viper, 프리미티드 타입의 포인터 변환이 용이한 k8 util외엔 추가했다고 생각합니다.

### 데이터 베이스 설계

#### 초기 데이터베이스 설계

유저(사장님)
- 휴대폰 번호와 비밀번호를 통해 회원 가입 할 수 있다.
카테고리
- 상품의 카테고리는 부모 카테고리 아이디를 갖으므로 계층적 구조를 갖을 수 있다.
상품
- 필수사항외에 초성검색을 위한 별도의 필드를 갖는다.
 

![](https://velog.velcdn.com/images/jakdangers/post/88d245ad-7e6d-4939-9a86-2f8a745ec5d8/image.png)

#### 최종 제출 (테이블이 변한 이유)

1. 카테고리 테이블을 삭제하게 된 이유는 일반적으로 카테고리는 계층적 구조를 갖기 때문에 확장성을 고려해 별도의 테이블로 생성하였으나 요구사항에 추가적인 정보가 없어서 오버엔지니어링 인듯 싶어 삭제했습니다.
2. 요구사항 중에 로그아웃을 할 수 있다는 것이 있는데 jwt 토큰 자체는 한번 발급하면 만료기한까지 핸들링 할 수 없는데 jwt을 세션처럼 데이터베이스에 저장하고 검증하도록 하기 위해 추가했습니다.
3. 이번에 초성 검색에 대해 처음 접했는데 고민이 많았습니다. 검색 레퍼런스도 적었고 오래 된 레퍼런스가 많았습니다. 대부분 RLIKE 쿼리로 처리하는게 많았습니다. 하지만 저는 저장하는 방식을 선택했습니다. 상품수가 많아지면 실시간으로 정규표현식 매칭하는게 데이터베이스 cpu사용량에 영향을 줄 것으로 예상되어 초성을 추출해 별도의 저장공간에 저장하는게 검색이나 성능에서 유리 할 것이고 비지니스 로직을 쿼리에 두지 않으려고 했습니다.

아래의 링크가 판단하는데 가장 큰 영향을 주었습니다.

[mysql rlike 커뮤니티](https://wetoz.kr/html/board.php?bo_table=tipntech&wr_id=67&sca=SQL&sfl=mb_id%2C1&stx=admin&sst=wr_nogood&sod=desc&sop=and&page=1)
[mysql 공식문서](https://dev.mysql.com/doc/refman/5.7/en/regexp.html)

![](https://velog.velcdn.com/images/jakdangers/post/8bb23dc7-9de9-404a-a2a1-23110608794a/image.png)

### 유저 도메인

유저 도메인은 휴대폰 번호와 패스워드 딱 필드 두개인데 생각보다 오래 걸렸습니다.

우선 휴대폰 번호의 포맷이 010-1234-5678 총 11자리가 되는 것만 고려하는 것인지 검색하는데 시간을 사용 했습니다. 요구사항에 전화번호 자리수가 명시되어 있지 않아 01X 또는 11자리 미만의 휴대폰 번호도 처리 해야하는지 관련 자료를 찾아보는데 시간을 사용했습니다.
그래서 최종적으로 휴대폰 번호는 11자리로 고정 된다는 것을 확인 했습니다.
휴대폰 번호의 입력값 또한 하이픈을 포함하는지 숫자로만 구성되어있는지 요구사항에 나와있지 않아 두 경우를 모두 유효한 값으로 가정하고 하이픈이 요청값에 있는 경우 정규표현식으로 제거했습니다.

패스워드의 경우 안전하게 보관 할 장치를 만들어야하는 요구사항이 있어 평문으로 저장하지 않고 암호화해서 저장 하도록 했습니다. 솔트 값은 높을수록 보안에 좋다고하나 리소스를 많이 사용해서 디폴트 값을 사용 했습니다.

### 상품 도메인

상품 도메인의 경우 필수 항목은 있었는데 별도의 제약 조건이 주어지지 않아서 어떤 점을 고려해야 할지 여러번 생각이 바뀌었던것 같습니다.
우선 카테고리부터 별도의 테이블이이었으나 요구사항에 따르면 오버엔지니어링 인 것 같아 컬럼화 했습니다.
가격, 원가의 경우 0원이 될수도 있고 소수점 계산이 필요 할수도 있을 것 같아 데시멀을 사용했습니다. 바코드의 경우 바코드는 무엇을 저장하지? 고민했었는데 바코드숫자 체계가 있어 디비에 50자 이내에 저장하면 될 것 이라고 생각했습니다.

### 토큰 도메인

유저의 로그아웃을 서버에서 핸들링 하기 위해 토큰스트링과 유효기간을 별도로 저장했습니다.
로그아웃 된 토큰을 사용한 요청이 오면 토큰의 활성 상태를 확인해 토큰의 유효기간이 만료되지 않았어도 만료된 것 처럼 사용했습니다.

### API 별 구현

#### 에러 응답 처리
GIN 바인딩 에러의 경우 커스텀 응답 포맷에 에러 메시지를 그대로 사용했고 그 외 비지니스로직에서 발생하는 에러의 경우 클라이언트 개발자가 에러 상황을 인지하고 대응 할 부분은 상세하게 기술 하였고 그 외 서버에서 처리 해야하는 부분은 서버에러가 발생했다고 하고 감추었습니다.

#### 유저
- CREATE USER - 패스워드는 별도의 제약 조건이 없어서 1자이상 255이하의 영어, 특수문자, 숫자 중 한글자를 포함하면 유효하다고 가정했습니다. 그리고 휴대폰 번호는 하이픈이 있는 형태와 없는 형태 두가지의 입력값만 유효하고 나머진 잘못 된 요청으로 처리했습니다.
컨트롤러와 서비스 계층에서 두번 검증하도록 했습니다. 현업에서는 두 계층을 다른 사람이 맡아서 구현 할 수 있기 때문에 컨트롤러에서 올바르게 입력값을 검증에서 온다고 가정하면 버그가 발생 할 수도 있기 때문입니다.
- LOGIN USER - 입력값의 올바른 포맷인지 확인하는데 집중했습니다.
- LOGOUT USER - JWT 미들웨어에서 어떤 메시지를 줄지 고민이 되었습니다.

#### 상품

- CREATE PRODUCT - 상품을 생성 할 때 추가적으로 초성을 추출한 것을 추가적으로 저장해야되서 한글 자음 추출 하는데 로직을 만들고 테스트 코드를 작성하는데 시간이 좀 걸렸습니다. 아스키코드처럼 유니코드의 크기 비교를 통해 한글문자가 아닌 경우 그대로 저장하고 한글 문자인 경우 자음을 추출했습니다.

- PATCH PRODUCT - uri 경로에 아이디를 넣어 구성 할지 고민이 되었는데 최종적으로 빼게 되었습니다. 바디를 사용 할 수 있는 경우 바디에 포함하는게 요청 정보가 분산되지 않고 클라이언트에서 사용하기 좋을것 같다고 생각했습니다.

- LIST PRODUCT - 한글 초성 검색을 위해 검색 키워드가 초성 그 외 문자로 이뤄진 경우와 한글문자를 포함하는 경우를 구분해
name 필드로 조회 할지 initial 필드로 조회할지 분기해 검색 하도록 했습니다.

- DELETE PRODUCT - 상품 삭제의 경우 소프트 딜리트, 하드 딜리트를 할지 고민 되었으나 데이터의 히스토리 상 소프트하게 지우는 방식으로 하는게 좋다고 생각했습니다.