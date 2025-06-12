// Collection page interactivity

const listsKey = 'brickholder_lists';
let lists = JSON.parse(localStorage.getItem(listsKey) || '[]');
lists.forEach(l => l.sets.forEach(s => { if (!s.quantity) s.quantity = 1; }));
let currentListId = null;

function saveLists() {
  localStorage.setItem(listsKey, JSON.stringify(lists));
}

function renderLists() {
  const panel = document.querySelector('.lists-panel');
  const listContainer = panel.querySelector('ul') || document.createElement('ul');
  listContainer.innerHTML = '';
  listContainer.className = 'list-items';
  lists.forEach(l => {
    const li = document.createElement('li');
    li.textContent = l.name;
    li.dataset.id = l.id;
    li.addEventListener('click', () => selectList(l.id));
    li.addEventListener('dblclick', () => {
      selectList(l.id);
      document.getElementById('list-form').style.display = 'block';
    });
    listContainer.appendChild(li);
  });
  if (!panel.contains(listContainer)) {
    const btn = panel.querySelector('.new-list-btn');
    if (btn) {
      panel.insertBefore(listContainer, btn);
    } else {
      panel.appendChild(listContainer);
    }
  }
}

function showEmptyState(show) {
  document.querySelector('.empty-state').style.display = show ? 'block' : 'none';
  document.querySelector('.list-section').style.display = 'none';
  document.querySelector('.lists-overview').style.display = 'none';
}

function showOverview() {
  renderListCards();
  document.querySelector('.empty-state').style.display = 'none';
  document.querySelector('.list-section').style.display = 'none';
  document.querySelector('.lists-overview').style.display = 'grid';
}

function selectList(id) {
  currentListId = id;
  const list = lists.find(l => l.id === id);
  if (!list) return;
  document.getElementById('list-title').textContent = list.name;
  document.getElementById('list-name').value = list.name;
  document.getElementById('list-desc').value = list.description || '';
  renderSetCards(list);
  document.querySelector('.empty-state').style.display = 'none';
  document.querySelector('.lists-overview').style.display = 'none';
  document.querySelector('.list-section').style.display = 'block';
}

function renderSetCards(list) {
  const grid = document.querySelector('.sets-grid');
  grid.innerHTML = '';
  list.sets.forEach(s => {
    const card = document.createElement('div');
    card.className = 'set-card';
    card.innerHTML = `
      <img src="${s.set_img_url || '../assets/sets.jpg'}" alt="${s.set_num}">
      <div class="set-info">
        <div class="set-code">${s.set_num}</div>
        <div class="set-name">${s.name}</div>
        <div class="set-parts">(${s.num_parts} деталей)</div>
        <div class="set-year">${s.year}</div>
      </div>
      <div class="set-qty">
        <button class="qty-minus">−</button>
        <span class="qty-value">${s.quantity}</span>
        <button class="qty-plus">+</button>
      </div>`;
    card.querySelector('.qty-minus').addEventListener('click', () => changeQuantity(list.id, s.set_num, -1));
    card.querySelector('.qty-plus').addEventListener('click', () => changeQuantity(list.id, s.set_num, 1));
    grid.appendChild(card);
  });
}

function renderListCards() {
  const grid = document.querySelector('.lists-overview');
  if (!grid) return;
  grid.innerHTML = '';
  lists.forEach(l => {
    const card = document.createElement('div');
    card.className = 'list-card';
    const first = l.sets[0];
    const imgSrc = first ? (first.set_img_url || '../assets/sets.jpg') : '../assets/sets.jpg';
    card.innerHTML = `
      <img src="${imgSrc}" alt="${l.name}">
      <div class="list-card-name">${l.name}</div>
    `;
    card.addEventListener('click', () => selectList(l.id));
    grid.appendChild(card);
  });
}

function changeQuantity(listId, setNum, delta) {
  const list = lists.find(l => l.id === listId);
  if (!list) return;
  const index = list.sets.findIndex(s => s.set_num === setNum);
  if (index === -1) return;
  const set = list.sets[index];
  set.quantity = (set.quantity || 1) + delta;
  if (set.quantity < 1) {
    list.sets.splice(index, 1);
  }
  saveLists();
  renderSetCards(list);
}

function newList() {
  const id = 'list-' + Date.now();
  const list = { id, name: 'Новый список', description: '', sets: [] };
  lists.push(list);
  saveLists();
  renderLists();
  selectList(id);
  document.getElementById('list-form').style.display = 'block';
}

function deleteList() {
  if (!currentListId) return;
  lists = lists.filter(l => l.id !== currentListId);
  saveLists();
  renderLists();
  if (lists.length === 0) {
    showEmptyState(true);
  } else {
    currentListId = null;
    showOverview();
  }
}

function saveCurrentList(e) {
  e.preventDefault();
  if (!currentListId) return;
  const list = lists.find(l => l.id === currentListId);
  list.name = document.getElementById('list-name').value.trim() || 'Без названия';
  list.description = document.getElementById('list-desc').value.trim();
  saveLists();
  renderLists();
  document.getElementById('list-title').textContent = list.name;
  document.getElementById('list-form').style.display = 'none';
}

function editList() {
  document.getElementById('list-form').style.display = 'block';
}

let searchTimeout;
async function fetchSearchResults() {
  const query = document.getElementById('search-input-set').value.trim();
  const container = document.getElementById('search-results');
  if (!query) {
    container.innerHTML = '';
    container.style.display = 'none';
    return;
  }

  const baseURL = 'http://localhost:8081';
  const res = await fetch(
    `${baseURL}/api/lego/sets?q=${encodeURIComponent(query)}`
  );
  if (!res.ok) return;
  const data = await res.json();
  const results = data.data || [];
  container.innerHTML = '';
  if (results.length === 0) {
    const div = document.createElement('div');
    div.className = 'search-item no-results';
    div.textContent = 'Ничего не найдено';
    container.appendChild(div);
  } else {
    results.forEach(s => {
      const div = document.createElement('div');
      div.className = 'search-item';
      div.innerHTML = `
        <img src="${s.set_img_url || '../assets/sets.jpg'}" alt="${s.set_num}">
        <span class="set-num">${s.set_num}</span>
        <span class="set-name">${s.name}</span>
      `;
      div.addEventListener('click', () => {
        addSetToCurrent(s);
        container.innerHTML = '';
        document.getElementById('search-input-set').value = '';
        container.style.display = 'none';
      });
      container.appendChild(div);
    });
  }
  container.style.display = 'block';
}

function searchSets() {
  clearTimeout(searchTimeout);
  searchTimeout = setTimeout(fetchSearchResults, 300);
}

function addSetToCurrent(set) {
  if (!currentListId) return;
  const list = lists.find(l => l.id === currentListId);
  if (list.sets.find(s => s.set_num === set.set_num)) return;
  list.sets.push({ ...set, quantity: 1 });
  saveLists();
  renderSetCards(list);
}

// header login display
function initHeader() {
  const loginLink = document.getElementById('login-link');
  const username = localStorage.getItem('username');
  if (username) {
    loginLink.textContent = username;
    loginLink.removeAttribute('href');
    loginLink.style.cursor = 'pointer';
    loginLink.addEventListener('click', (e) => {
      e.preventDefault();
      localStorage.removeItem('username');
      window.location.href = 'index.html';
    });
  } else {
    loginLink.textContent = 'Войти';
    loginLink.setAttribute('href', 'login.html');
  }
}

document.addEventListener('DOMContentLoaded', () => {
  renderLists();
  if (lists.length === 0) {
    showEmptyState(true);
  } else {
    showOverview();
  }
  document.querySelector('.new-list-btn').addEventListener('click', newList);
  document.getElementById('delete-btn').addEventListener('click', deleteList);
  document.getElementById('list-form').addEventListener('submit', saveCurrentList);
  document.getElementById('edit-btn').addEventListener('click', editList);
  document.getElementById('add-set-btn').addEventListener('click', fetchSearchResults);
  const searchInputSet = document.getElementById('search-input-set');
  const resultsContainer = document.getElementById('search-results');
  resultsContainer.style.display = 'none';
  searchInputSet.addEventListener('input', searchSets);
  searchInputSet.addEventListener('focus', () => {
    if (searchInputSet.value.trim()) {
      fetchSearchResults();
    } else if (resultsContainer.innerHTML) {
      resultsContainer.style.display = 'block';
    }
  });
  document.addEventListener('click', (e) => {
    if (!resultsContainer.contains(e.target) && e.target !== searchInputSet) {
      resultsContainer.style.display = 'none';
    }
  });
  initHeader();
});
