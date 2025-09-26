// Vercel Node.js Function
export default function handler(req, res) {
  // Set CORS headers
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');

  // Handle preflight requests
  if (req.method === 'OPTIONS') {
    res.status(200).end();
    return;
  }

  const { url } = req;

  // Route handling
  if (url === '/health') {
    return handleHealth(req, res);
  }

  if (url === '/api/v1/users') {
    return handleUsers(req, res);
  }

  if (url === '/api/v1/projects') {
    return handleProjects(req, res);
  }

  if (url === '/api/v1/skills') {
    return handleSkills(req, res);
  }

  // Default 404
  res.status(404).json({
    error: 'Endpoint not found',
    path: url
  });
}

function handleHealth(req, res) {
  res.status(200).json({
    status: 'healthy',
    timestamp: new Date().toISOString(),
    version: '1.0.0',
    platform: 'vercel-nodejs'
  });
}

function handleUsers(req, res) {
  const users = [
    {
      id: '1',
      name: '이혁주',
      email: 'hyoukjoo@example.com',
      role: 'Full-Stack Developer',
      bio: 'Backend engineer specializing in Go, TypeScript, and cloud architecture',
      website: 'https://hyoukjoolee.github.io/portfolio',
      location: 'Seoul, Korea',
      skills: ['Go', 'TypeScript', 'Flutter', 'AWS', 'Docker', 'PostgreSQL'],
      is_public: true
    }
  ];

  res.status(200).json({
    data: users,
    count: users.length
  });
}

function handleProjects(req, res) {
  const projects = [
    {
      id: '1',
      title: 'Portfolio API Server',
      description: 'Go REST API with Supabase integration deployed on Vercel',
      tech_stack: ['Go', 'Supabase', 'Vercel'],
      status: 'completed',
      featured: true,
      live_url: 'https://portfolio-api.vercel.app',
      github_url: 'https://github.com/hyoukjoolee/portfolio-api',
      is_public: true
    },
    {
      id: '2',
      title: 'Flutter Portfolio',
      description: 'Responsive portfolio website with i18n support',
      tech_stack: ['Flutter', 'Dart', 'GitHub Pages'],
      status: 'completed',
      featured: true,
      live_url: 'https://hyoukjoolee.github.io/portfolio',
      github_url: 'https://github.com/hyoukjoolee/portfolio',
      is_public: true
    }
  ];

  res.status(200).json({
    data: projects,
    count: projects.length
  });
}

function handleSkills(req, res) {
  const skills = [
    { id: '1', name: 'Go', category: 'backend', level: 'expert', years_exp: 3, featured: true, color: '#00ADD8' },
    { id: '2', name: 'TypeScript', category: 'frontend', level: 'advanced', years_exp: 4, featured: true, color: '#3178C6' },
    { id: '3', name: 'Flutter', category: 'mobile', level: 'advanced', years_exp: 2, featured: true, color: '#02569B' },
    { id: '4', name: 'Supabase', category: 'database', level: 'intermediate', years_exp: 1, featured: true, color: '#3ECF8E' }
  ];

  res.status(200).json({
    data: skills,
    count: skills.length
  });
}