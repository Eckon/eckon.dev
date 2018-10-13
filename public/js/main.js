// main nav redirect on click
(function () {
    let items = document.querySelector('#main-nav').querySelectorAll('div');
    items.forEach((ele) => {
        ele.addEventListener('click', () => {
            window.location.href = ele.dataset.path;
        })
    })
})();
