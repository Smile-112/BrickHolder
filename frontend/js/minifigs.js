
document.addEventListener('DOMContentLoaded', () => {
  const figsTree = document.getElementById('figs-tree');
  const modal = document.getElementById('fig-modal');
  const modalBody = document.getElementById('modal-body');
  const modalClose = document.getElementById('modal-close');

  function showModal(fig) {
    modalBody.innerHTML = `
      <h3>${fig.name}</h3>
      <img src="${fig.set_img_url}" alt="${fig.name}">
      <p>Артикул: ${fig.set_num}</p>
      <p>Количество деталей: ${fig.num_parts}</p>
      <a href="${fig.set_url}" target="_blank">Ссылка на Rebrickable</a>
    `;
    modal.classList.add('open');
  }

  modalClose.addEventListener('click', () => modal.classList.remove('open'));
  modal.addEventListener('click', e => { if (e.target === modal) modal.classList.remove('open'); });

  fetch('http://localhost:8081/api/lego/minifigs')
    .then(r => r.json())
    .then(d => d.data || [])
    .then(figs => {
      const frag = document.createDocumentFragment();
      figs.forEach(fig => {
        const li = document.createElement('li');
        li.className = 'set-item';
        const img = document.createElement('img');
        img.src = fig.set_img_url;
        li.appendChild(img);
        const span = document.createElement('span');
        span.textContent = fig.name;
        li.appendChild(span);
        li.addEventListener('click', () => showModal(fig));
        frag.appendChild(li);
      });
      figsTree.innerHTML = '';
      figsTree.appendChild(frag);
    })
    .catch(() => { figsTree.textContent = 'Ошибка загрузки данных'; });
});

