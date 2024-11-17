const hamburger = document.querySelector('#hamburger');
const nav = document.querySelector('nav');

hamburger.addEventListener('click', () => {
    hamburger.classList.toggle('open');
    nav.classList.toggle('open');

    document.body.classList.toggle("frozen");
});

document.addEventListener('click', (event) => {
    console.log('user clicked on:', event.target);
});