# 🆓 **100% 무료 배포 가이드**

포트폴리오 API를 **완전 무료**로 배포하는 방법들을 **Docker/Non-Docker**로 구분해서 제공합니다.

---

## 🚀 **Non-Docker 배포 (추천!)**

### 🥇 **방법 1: Vercel (가장 간단)**

#### 장점
- ✅ **완전 무료**: Hobby 플랜 평생 무료
- ✅ **Go 네이티브**: 별도 설정 없음
- ✅ **즉시 배포**: 1분 만에 완료
- ✅ **글로벌 CDN**: 전세계 빠른 속도
- ✅ **자동 HTTPS**: SSL 자동

#### 배포 단계
1. **Vercel CLI 설치**:
   ```bash
   npm i -g vercel
   ```

2. **배포**:
   ```bash
   vercel
   ```

3. **환경변수 설정**:
   ```bash
   vercel env add SUPABASE_URL
   vercel env add SUPABASE_ANON_KEY
   vercel env add SUPABASE_DB_URL
   ```

**예상 URL**: `https://portfolio-api.vercel.app`

---

### 🥈 **방법 2: Netlify Functions**

#### 장점
- ✅ **125K 요청/월** 무료
- ✅ **서버리스**: 인프라 관리 불필요
- ✅ **Git 연동**: 자동 배포

#### 배포 단계
1. **Netlify에 Git 연결**
2. **환경변수 설정**
3. **자동 빌드 설정 완료**

**예상 URL**: `https://portfolio-api.netlify.app`

---

## 🐳 **Docker 배포**

### 🥇 **방법 1: Railway**

#### 장점
- ✅ **$5 크레딧/월**: 충분한 무료 사용량
- ✅ **Docker 지원**: 기존 Dockerfile 그대로
- ✅ **PostgreSQL**: 필요시 추가 DB

#### 배포 단계
1. **Railway 회원가입**: https://railway.app
2. **GitHub 연결** 후 저장소 선택
3. **환경변수 설정**:
   ```
   SUPABASE_URL=https://your-project.supabase.co
   SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   SUPABASE_DB_URL=postgresql://postgres:password@db.project.supabase.co:5432/postgres
   GIN_MODE=release
   ```
4. **자동 배포**: Docker 이미지 빌드

**예상 URL**: `https://portfolio-api.up.railway.app`

---

### 🥈 **방법 2: Render**

#### 장점
- ✅ **750시간/월** 무료
- ✅ **자동 HTTPS**
- ❌ **15분 후 슬립** (첫 요청 시 지연)

#### 배포 단계
1. **Render 회원가입**: https://render.com
2. **Web Service 생성**:
   - Repository 연결
   - Environment: Docker
   - Plan: Free
3. **환경변수 설정** (Railway와 동일)

**예상 URL**: `https://portfolio-api.onrender.com`

---

### 🥉 **방법 3: Fly.io**

#### 장점
- ✅ **3GB RAM, 1000시간/월**
- ✅ **글로벌 엣지**
- ✅ **빠른 성능**

#### 배포 단계
1. **Fly.io CLI 설치**:
   ```bash
   # macOS
   curl -L https://fly.io/install.sh | sh
   ```

2. **앱 생성 & 배포**:
   ```bash
   flyctl launch
   flyctl deploy
   ```

**예상 URL**: `https://portfolio-api.fly.dev`

---

## 📊 **무료 플랜 비교**

### Non-Docker (서버리스)
| 서비스 | 요청 한도 | 콜드 스타트 | 복잡도 | 추천도 |
|--------|----------|------------|--------|--------|
| **Vercel** | Fair Use | ~100ms | ⭐ | 🥇 |
| **Netlify** | 125K/월 | ~200ms | ⭐⭐ | 🥈 |

### Docker (컨테이너)
| 서비스 | 실행시간 | 슬립모드 | 복잡도 | 추천도 |
|--------|----------|----------|--------|--------|
| **Railway** | 500시간 | ❌ 없음 | ⭐⭐⭐ | 🥇 |
| **Render** | 750시간 | ✅ 15분 | ⭐⭐ | 🥈 |
| **Fly.io** | 1000시간 | ❌ 없음 | ⭐⭐⭐⭐ | 🥉 |

## 🎯 **추천 순서**

### 🚀 **간단함 우선**
1. **Vercel** (Non-Docker) - 1분 배포
2. **Netlify** (Non-Docker) - 서버리스
3. **Railway** (Docker) - 기존 코드 재사용

### 🏃 **성능 우선**
1. **Fly.io** (Docker) - 글로벌 엣지
2. **Railway** (Docker) - 안정적
3. **Vercel** (Non-Docker) - CDN 빠름

## 💡 **결론**
- **처음 시작**: Vercel (가장 쉬움)
- **Docker 유지**: Railway (설정 최소)
- **확장 계획**: Fly.io (프로덕션급)

## 🔧 **공통 환경변수**

모든 플랫폼에서 동일하게 설정:
```env
GIN_MODE=release
PORT=8080
DATABASE_URL=postgresql://username:password@host:port/database
```

## 🔗 **배포 후 확인**

배포 완료 후 헬스체크:
```bash
curl https://your-app-url.com/health
```

예상 응답:
```json
{
  "status": "healthy",
  "timestamp": "2024-09-26T...",
  "version": "1.0.0"
}
```

## 💡 **비용 최적화 팁**

1. **리소스 최소화**: 1 CPU, 512MB RAM
2. **슬립 모드 활용**: 사용하지 않을 때 자동 슬립
3. **로그 제한**: CloudWatch 로그 최소화
4. **이미지 최적화**: 멀티스테이지 Docker 빌드

---

## 🆘 **문제 해결**

### Railway 배포 실패
```bash
# 로그 확인
railway logs

# 재배포
railway up
```

### Render 슬립 해제
```bash
# 웜업 요청 (cron job 설정 가능)
curl https://your-app.onrender.com/health
```

### Fly.io 문제
```bash
# 로그 확인
flyctl logs

# 상태 확인
flyctl status
```

---

**모든 옵션이 100% 무료입니다!** 🎉
가장 편한 방법을 선택해서 배포하세요!