document.addEventListener('DOMContentLoaded', () => {
  const seriesTree = document.getElementById('series-tree');
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
    series.forEach(s => {
      const pid = s.parent_id || 0;
      if (!children[pid]) children[pid] = [];
      children[pid].push(s);
    });

    function buildSeries(sr) {
      const li = document.createElement('li');
      const header = document.createElement('div');
      header.className = 'series-item';
      const img = document.createElement('img');
      const list = setsBySeries[sr.id] || [];
      img.src = list.length ? list[0].set_img_url : '../assets/sets.jpg';
      header.appendChild(img);
      const span = document.createElement('span');
      span.textContent = sr.name;
      header.appendChild(span);
      li.appendChild(header);

      const childContainer = document.createElement('ul');
      childContainer.className = 'collapsible';
      li.appendChild(childContainer);

      header.addEventListener('click', () => {
        childContainer.classList.toggle('open');
      });

      if (children[sr.id]) {
        children[sr.id].forEach(ch => childContainer.appendChild(buildSeries(ch)));
      }

      list.forEach(set => childContainer.appendChild(buildSet(set)));

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

    (children[0] || []).forEach(sr => seriesTree.appendChild(buildSeries(sr)));
  }).catch(() => {
    seriesTree.textContent = 'Ошибка загрузки данных';
  });
});
