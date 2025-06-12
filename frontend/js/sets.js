document.addEventListener('DOMContentLoaded', () => {
  const seriesTree = document.getElementById('series-tree');
  const breadcrumb = document.getElementById('breadcrumb');
  const modal = document.getElementById('set-modal');
  const modalBody = document.getElementById('modal-body');
  const modalClose = document.getElementById('modal-close');

  modalClose.addEventListener('click', () => modal.classList.remove('open'));
  modal.addEventListener('click', e => { if (e.target === modal) modal.classList.remove('open'); });

  Promise.all([
    fetch('http://localhost:8081/api/lego/series').then(r => r.json()).then(d => d.data || []),
    fetch('http://localhost:8081/api/lego/sets').then(r => r.json()).then(d => d.data || [])
  ]).then(([series, sets]) => {
    const setsBySeries = {};
    sets.forEach(s => {
      (setsBySeries[s.theme_id] = setsBySeries[s.theme_id] || []).push(s);
    });

    const children = {};
    const seriesMap = {};
    series.forEach(s => {
      seriesMap[s.id] = s;
      const pid = s.parent_id || 0;
      if (!children[pid]) children[pid] = [];
      children[pid].push(s);
    });

    function buildSeries(sr) {
      const li = document.createElement('li');
      li.className = 'series-item';
      const img = document.createElement('img');
      const list = setsBySeries[sr.id] || [];
      img.src = list.length ? list[0].set_img_url : '../assets/sets.jpg';
      li.appendChild(img);
      const span = document.createElement('span');
      span.textContent = sr.name;
      li.appendChild(span);
      li.addEventListener('click', () => renderTree(sr.id));
      return li;
    }

    function buildSet(set) {
      const li = document.createElement('li');
      li.className = 'set-item';
      const img = document.createElement('img');
      img.src = set.set_img_url;
      li.appendChild(img);
      const span = document.createElement('span');
      span.textContent = set.name;
      li.appendChild(span);
      li.addEventListener('click', e => {
        e.stopPropagation();
        showModal(set);
      });
      return li;
    }

    function showModal(set) {
      modalBody.innerHTML = `
        <h3>${set.name}</h3>
        <img src="${set.set_img_url}" alt="${set.name}">
        <p>Артикул: ${set.set_num}</p>
        <p>Год выпуска: ${set.year}</p>
        <p>Количество деталей: ${set.num_parts}</p>
        <a href="${set.set_url}" target="_blank">Ссылка на Rebrickable</a>
      `;
      modal.classList.add('open');
    }

    function buildBreadcrumb(id) {
      breadcrumb.innerHTML = '';
      const path = [];
      let cur = seriesMap[id];
      while (cur) {
        path.unshift(cur);
        cur = cur.parent_id ? seriesMap[cur.parent_id] : null;
      }
      const rootLink = document.createElement('a');
      rootLink.href = '#';
      rootLink.textContent = 'Наборы';
      rootLink.addEventListener('click', e => { e.preventDefault(); renderTree(0); });
      breadcrumb.appendChild(rootLink);
      path.forEach(sr => {
        breadcrumb.appendChild(document.createTextNode(' / '));
        const a = document.createElement('a');
        a.href = '#';
        a.textContent = sr.name;
        a.addEventListener('click', e => { e.preventDefault(); renderTree(sr.id); });
        breadcrumb.appendChild(a);
      });
    }

    function renderTree(rootId) {
      const frag = document.createDocumentFragment();
      (children[rootId] || []).forEach(sr => frag.appendChild(buildSeries(sr)));
      (setsBySeries[rootId] || []).forEach(set => frag.appendChild(buildSet(set)));
      seriesTree.innerHTML = '';
      seriesTree.appendChild(frag);
      buildBreadcrumb(rootId);
      seriesTree.classList.add('slide-enter');
      requestAnimationFrame(() => seriesTree.classList.add('slide-enter-active'));
      seriesTree.addEventListener('transitionend', () => {
        seriesTree.classList.remove('slide-enter','slide-enter-active');
      }, { once: true });
    }

    renderTree(0);
  }).catch(() => {
    seriesTree.textContent = 'Ошибка загрузки данных';
  });
});
