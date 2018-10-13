// main nav redirect on click
(function () {
    let items = document.querySelector('#main-nav').querySelectorAll('div');
    items.forEach((ele) => {
        ele.parentNode.addEventListener('click', () => {
            window.location.href = ele.dataset.path;
        })
    })
})();

// show the vertical nav bar when click action on small-nav-header (when screen is small - mobile)
(function () {
    let nav = document.querySelector('#small-nav-header');
    nav.addEventListener('click', () => {
        const s = document.querySelector('#main-nav').style;
        s.display = s.display === 'block' ? 'none' : 'block';
    })
})();
