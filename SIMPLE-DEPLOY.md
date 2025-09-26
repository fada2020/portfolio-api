# 🚀 **초간단 무료 배포 가이드**

Supabase와 함께하는 **오버스펙 없는** 심플한 무료 배포!

## 🎯 **왜 간단한가?**

- ✅ **데이터베이스**: Supabase가 모든 걸 해결
- ✅ **인증**: Supabase 내장
- ✅ **실시간**: Supabase WebSocket
- ✅ **스토리지**: Supabase 파일 업로드
- ❌ **복잡한 설정 없음**: Docker, Kubernetes, 로드밸런서 필요없음

---

## 🥇 **방법 1: Vercel (가장 간단)**

### 장점
- ✅ **완전 무료**: Hobby 플랜 평생 무료
- ✅ **Go 지원**: 네이티브 Go 지원
- ✅ **즉시 배포**: Git push만으로 배포
- ✅ **글로벌 CDN**: 전세계 빠른 속도
- ✅ **자동 HTTPS**: SSL 인증서 자동

### 배포 단계
1. **vercel.json 생성**:
   ```json
   {
     "functions": {
       "api/index.go": {
         "runtime": "@vercel/go"
       }
     }
   }
   ```

2. **main.go를 api/index.go로 복사**

3. **Vercel CLI 설치**:
   ```bash
   npm i -g vercel
   ```

4. **배포**:
   ```bash
   vercel
   ```

5. **환경변수 설정**:
   ```bash
   vercel env add SUPABASE_URL
   vercel env add SUPABASE_ANON_KEY
   vercel env add SUPABASE_DB_URL
   ```

**예상 URL**: `https://portfolio-api.vercel.app`

---

## 🥈 **방법 2: Netlify Functions**

### 장점
- ✅ **완전 무료**: 125K 요청/월
- ✅ **서버리스**: 인프라 관리 불필요
- ✅ **Git 연동**: 자동 배포

### 배포 단계
1. **netlify.toml 생성**:
   ```toml
   [build]
     functions = "functions"

   [[redirects]]
     from = "/api/*"
     to = "/.netlify/functions/:splat"
     status = 200
   ```

2. **functions 폴더 생성 후 Go 함수 작성**

3. **Netlify에 Git 연결**

**예상 URL**: `https://portfolio-api.netlify.app`

---

## 🥉 **방법 3: Railway (Docker 원한다면)**

### 장점
- ✅ **$5 크레딧/월**: 충분한 무료 사용량
- ✅ **Docker 지원**: 기존 Dockerfile 그대로 사용
- ✅ **PostgreSQL**: 필요시 추가 DB

### 배포 단계
1. **Railway 가입**: https://railway.app
2. **GitHub 연결**: 저장소 선택
3. **환경변수 설정**:
   - `SUPABASE_URL`
   - `SUPABASE_ANON_KEY`
   - `SUPABASE_DB_URL`
4. **자동 배포**: Docker 이미지 빌드

**예상 URL**: `https://portfolio-api.up.railway.app`

---

## 🏆 **추천 선택**

### 🎯 **최대한 간단하게**: Vercel
- 설정 파일 1개만 만들면 끝
- Go 코드 그대로 업로드
- 무료 + 빠름 + 안정적

### ⚡ **기존 Docker 유지**: Railway
- Dockerfile 그대로 사용 가능
- 추가 설정 최소화
- PostgreSQL 확장 가능

### 🔧 **커스터마이징**: Netlify
- 서버리스 함수 세밀 제어
- JAMstack 생태계 활용

---

## 📊 **비교표**

| 서비스 | 복잡도 | 무료 한도 | Go 지원 | 배포 시간 |
|--------|-------|----------|---------|----------|
| **Vercel** | ⭐ | 무제한* | ✅ 네이티브 | 1분 |
| **Netlify** | ⭐⭐ | 125K/월 | ✅ 함수 | 2분 |
| **Railway** | ⭐⭐⭐ | $5/월 | ✅ Docker | 3분 |

*Fair Use Policy 적용

---

## 🔧 **공통 환경변수**

모든 플랫폼에서 동일하게 설정:

```env
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
SUPABASE_DB_URL=postgresql://postgres:password@db.project.supabase.co:5432/postgres
GIN_MODE=release
```

---

## 💡 **결론**

**Railway 오버스펙 맞습니다!**

단순한 Go API라면:
1. **Vercel** - 가장 간단
2. **Netlify** - 서버리스 선호시
3. **Railway** - Docker 재사용시

**Supabase만으로도 충분합니다!** 🎉