# 🚀 **Vercel 배포 완전 가이드**

**5분 안에** Supabase + Go API를 Vercel로 완전 무료 배포하기!

---

## 📋 **Step 1: GitHub 저장소 생성**

### 1.1 GitHub에서 새 저장소 생성
1. **GitHub.com 접속** → **New repository**
2. **Repository name**: `portfolio-api`
3. **Public** 선택 (무료 배포용)
4. **Create repository** 클릭

### 1.2 로컬에서 Git 초기화
```bash
cd /path/to/portfolio-api

# Git 초기화
git init
git add .
git commit -m "🎉 Initial commit: Supabase + Go API"

# GitHub 연결 (YOUR_USERNAME을 본인 것으로 변경)
git remote add origin https://github.com/YOUR_USERNAME/portfolio-api.git
git branch -M main
git push -u origin main
```

---

## 📋 **Step 2: Vercel CLI 설치 및 로그인**

### 2.1 Vercel CLI 설치
```bash
# npm으로 설치
npm i -g vercel

# 또는 yarn으로 설치
yarn global add vercel
```

### 2.2 Vercel 로그인
```bash
vercel login
```
- **GitHub 계정**으로 로그인 (추천)
- 브라우저에서 인증 완료

---

## 📋 **Step 3: Supabase 프로젝트 생성**

### 3.1 Supabase 가입 및 프로젝트 생성
1. **https://supabase.com** 접속
2. **Start your project** 클릭
3. **GitHub로 로그인**
4. **New project** 생성:
   - **Organization**: Personal
   - **Project Name**: `portfolio-api`
   - **Database Password**: 강력한 비밀번호 생성
   - **Region**: `Northeast Asia (Seoul)`

### 3.2 환경변수 정보 수집
프로젝트 생성 후 **Settings → API**에서 확인:

```env
SUPABASE_URL=https://your-project-ref.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Settings → Database → Connection string**에서 확인:
```env
SUPABASE_DB_URL=postgresql://postgres:your-password@db.your-project-ref.supabase.co:5432/postgres
```

---

## 📋 **Step 4: Vercel 프로젝트 생성 및 배포**

### 4.1 Vercel 프로젝트 초기화
```bash
cd /path/to/portfolio-api
vercel
```

질문에 답변:
- **Set up and deploy?**: `Y`
- **Which scope?**: 본인 계정 선택
- **Link to existing project?**: `N`
- **Project name**: `portfolio-api` (엔터)
- **Directory**: `.` (엔터)
- **Override settings?**: `N`

### 4.2 배포 완료!
- **Production URL**: `https://portfolio-api-xxx.vercel.app`
- 약 **10초 후** 배포 완료!

---

## 📋 **Step 5: 환경변수 설정**

### 5.1 Vercel에 환경변수 추가
```bash
# Supabase URL 설정
vercel env add SUPABASE_URL

# Supabase Anonymous Key 설정
vercel env add SUPABASE_ANON_KEY

# Supabase Database URL 설정
vercel env add SUPABASE_DB_URL

# Gin Mode 설정
vercel env add GIN_MODE
```

각 변수에 대해:
- **Environment**: `Production, Preview, Development` 모두 선택
- **Value**: Step 3에서 수집한 값 입력

### 5.2 재배포
```bash
vercel --prod
```

---

## 📋 **Step 6: 배포 테스트**

### 6.1 Health Check 테스트
```bash
curl https://your-app.vercel.app/health
```

예상 응답:
```json
{
  "status": "healthy",
  "timestamp": "2024-09-26T10:00:00Z",
  "version": "1.0.0",
  "platform": "vercel"
}
```

### 6.2 API 엔드포인트 테스트
```bash
# 사용자 목록 조회
curl https://your-app.vercel.app/api/v1/users

# 프로젝트 목록 조회
curl https://your-app.vercel.app/api/v1/projects

# 기술 스택 조회
curl https://your-app.vercel.app/api/v1/skills
```

### 6.3 Swagger 문서 확인
브라우저에서 접속:
```
https://your-app.vercel.app/swagger/index.html
```

---

## 📋 **Step 7: 포트폴리오 연동**

### 7.1 OpenAPI 스펙 업데이트
포트폴리오의 `portfolio-api.json` 파일에서 host 업데이트:

```json
{
  "host": "your-app.vercel.app",
  "schemes": ["https"]
}
```

### 7.2 포트폴리오에서 테스트
1. 포트폴리오 빌드 및 배포
2. **API** 페이지에서 **Portfolio API** 선택
3. 실제 API 엔드포인트 테스트

---

## 🎉 **완료!**

### ✅ **성공 확인 체크리스트**
- [ ] GitHub 저장소 생성 및 코드 업로드
- [ ] Vercel 프로젝트 배포 성공
- [ ] Supabase 데이터베이스 연결 성공
- [ ] 환경변수 설정 완료
- [ ] Health check 200 응답
- [ ] API 엔드포인트 정상 동작
- [ ] Swagger 문서 접근 가능
- [ ] 포트폴리오에서 API 테스트 성공

### 🔗 **최종 결과**
- **API URL**: `https://your-app.vercel.app`
- **Swagger 문서**: `https://your-app.vercel.app/swagger/index.html`
- **비용**: **$0** (완전 무료!)
- **관리**: **자동** (Git push만으로 재배포)

---

## 🛠️ **향후 관리**

### 자동 배포
Git push만으로 자동 배포:
```bash
git add .
git commit -m "✨ Add new feature"
git push
```
→ Vercel이 자동으로 배포!

### 환경변수 관리
```bash
# 환경변수 목록 확인
vercel env ls

# 환경변수 제거
vercel env rm VARIABLE_NAME
```

### 로그 확인
```bash
# 실시간 로그 확인
vercel logs
```

---

**🎯 총 소요시간: 5분 이내**
**💰 총 비용: $0**
**🔧 관리 부담: 거의 없음**