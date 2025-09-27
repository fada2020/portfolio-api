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

  if (url === '/api/v1/contact') {
    return handleContact(req, res);
  }

  // Individual resource endpoints
  if (url.startsWith('/api/v1/users/') && url !== '/api/v1/users') {
    return handleUserById(req, res);
  }

  if (url.startsWith('/api/v1/projects/') && url !== '/api/v1/projects') {
    return handleProjectById(req, res);
  }

  if (url.startsWith('/api/v1/skills/') && url !== '/api/v1/skills') {
    return handleSkillById(req, res);
  }

  // Stats endpoints
  if (url === '/api/v1/stats/views') {
    return handleStatsViews(req, res);
  }

  if (url === '/api/v1/stats/projects') {
    return handleStatsProjects(req, res);
  }

  if (url === '/api/v1/stats/visit') {
    return handleStatsVisit(req, res);
  }

  // Root health check
  if (url === '/') {
    return handleHealth(req, res);
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

function handleContact(req, res) {
  if (req.method !== 'POST') {
    return res.status(405).json({
      error: 'Method not allowed',
      message: 'Only POST requests are supported for this endpoint'
    });
  }

  // In a real implementation, you would:
  // 1. Validate the request body
  // 2. Send email or store in database
  // 3. Return appropriate response

  const contactId = Math.floor(Math.random() * 10000);

  res.status(201).json({
    message: 'Contact form submitted successfully',
    id: contactId,
    timestamp: new Date().toISOString()
  });
}

function handleUserById(req, res) {
  const userId = req.url.split('/').pop();

  if (req.method === 'GET') {
    // Get single user
    res.status(200).json({
      id: userId,
      name: '이혁주',
      email: 'hyoukjoo@example.com',
      role: 'Full-Stack Developer',
      bio: 'Backend engineer specializing in Go, TypeScript, and cloud architecture',
      website: 'https://hyoukjoolee.github.io/portfolio',
      location: 'Seoul, Korea',
      skills: ['Go', 'TypeScript', 'Flutter', 'AWS', 'Docker', 'PostgreSQL'],
      is_public: true
    });
  } else if (req.method === 'PUT') {
    // Update user
    res.status(200).json({
      message: 'User updated successfully',
      id: userId,
      timestamp: new Date().toISOString()
    });
  } else if (req.method === 'DELETE') {
    // Delete user
    res.status(204).end();
  } else {
    res.status(405).json({ error: 'Method not allowed' });
  }
}

function handleProjectById(req, res) {
  const projectId = req.url.split('/').pop();

  if (req.method === 'GET') {
    // Get single project
    res.status(200).json({
      id: projectId,
      title: 'Portfolio API Server',
      description: 'Go REST API with Supabase integration deployed on Vercel',
      tech_stack: ['Go', 'Supabase', 'Vercel'],
      status: 'completed',
      featured: true,
      live_url: 'https://portfolio-api.vercel.app',
      github_url: 'https://github.com/hyoukjoolee/portfolio-api',
      is_public: true
    });
  } else if (req.method === 'PUT') {
    // Update project
    res.status(200).json({
      message: 'Project updated successfully',
      id: projectId,
      timestamp: new Date().toISOString()
    });
  } else if (req.method === 'DELETE') {
    // Delete project
    res.status(204).end();
  } else {
    res.status(405).json({ error: 'Method not allowed' });
  }
}

function handleSkillById(req, res) {
  const skillId = req.url.split('/').pop();

  if (req.method === 'DELETE') {
    // Delete skill
    res.status(204).end();
  } else {
    res.status(405).json({ error: 'Method not allowed' });
  }
}

function handleStatsViews(req, res) {
  res.status(200).json({
    total_views: 1250,
    unique_visitors: 950,
    views_today: 45,
    views_this_week: 320,
    views_this_month: 1100,
    top_pages: [
      { page: '/projects', views: 450 },
      { page: '/api', views: 320 },
      { page: '/resume', views: 280 }
    ],
    views_by_country: [
      { country: 'KR', views: 750 },
      { country: 'US', views: 250 },
      { country: 'JP', views: 150 }
    ],
    views_over_time: [
      { date: '2023-12-01', views: 45 },
      { date: '2023-12-02', views: 38 },
      { date: '2023-12-03', views: 52 }
    ]
  });
}

function handleStatsProjects(req, res) {
  res.status(200).json({
    total_projects: 12,
    completed_projects: 8,
    featured_projects: 3,
    tech_stack_stats: [
      { technology: 'Go', count: 5, percentage: 41.7 },
      { technology: 'TypeScript', count: 7, percentage: 58.3 },
      { technology: 'Flutter', count: 3, percentage: 25.0 }
    ],
    projects_by_status: [
      { status: 'completed', count: 8 },
      { status: 'in-progress', count: 3 },
      { status: 'planned', count: 1 }
    ]
  });
}

function handleStatsVisit(req, res) {
  if (req.method !== 'POST') {
    return res.status(405).json({
      error: 'Method not allowed',
      message: 'Only POST requests are supported for this endpoint'
    });
  }

  res.status(201).json({
    message: 'Visit recorded successfully',
    timestamp: new Date().toISOString()
  });
}