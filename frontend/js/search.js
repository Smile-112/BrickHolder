document.addEventListener('DOMContentLoaded', () => {
  const searchInput = document.getElementById('search-input');
  if (!searchInput) return;
  let searchTimeout;

  const container = document.createElement('div');
  container.className = 'search-results';
  container.style.position = 'absolute';
  container.style.top = '100%';
  container.style.left = '0';
  container.style.right = '0';
  container.style.display = 'none';

  const parent = searchInput.parentElement;
  parent.style.position = 'relative';
  parent.appendChild(container);

  async function fetchResults() {
    const query = searchInput.value.trim();
    if (!query) {
      container.innerHTML = '';
      container.style.display = 'none';
      return;
    }
    try {
      const res = await fetch(`http://localhost:8081/api/lego/sets?q=${encodeURIComponent(query)}`);
      if (!res.ok) return;
      const data = await res.json();
      const results = data.data || [];
      container.innerHTML = '';
      results.forEach(s => {
        const div = document.createElement('div');
        div.className = 'search-item';
        div.textContent = `${s.set_num} - ${s.name}`;
        container.appendChild(div);
      });
      container.style.display = results.length ? 'block' : 'none';
    } catch (_) {
      container.style.display = 'none';
    }
  }

  searchInput.addEventListener('input', () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(fetchResults, 300);
  });

  document.addEventListener('click', (e) => {
    if (!container.contains(e.target) && e.target !== searchInput) {
      container.style.display = 'none';
    }
  });
});
