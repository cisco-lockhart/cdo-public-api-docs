(function () {
    PubHub.util.loadJSFile('https://cdnjs.cloudflare.com/ajax/libs/slick-carousel/1.9.0/slick.js').then(function () {
        try{
            $('#slider').slick({
                infinite: true,
                speed: 300,
                slidesToShow: 3,
                slidesToScroll: 1,
                responsive: [
                    {
                        breakpoint: 992,
                        settings: {
                            slidesToShow: 2,
                            slidesToScroll: 1,
                            infinite: true,
                        }
                    },
                    {
                        breakpoint: 768,
                        settings: {
                            slidesToShow: 1,
                            slidesToScroll: 1,
                            infinite: true,
                        }
                    },
                    {
                        breakpoint: 576,
                        settings: {
                            slidesToShow: 1,
                            slidesToScroll: 1,
                            infinite: true,
                            variableWidth: true,
                        }
                    }
                ]
            });
            document.querySelector('.ph-impt-section').style.height = 'auto';
            document.querySelector('.ph-impt-section').style.opacity = '1';
        } catch(e){
            console.error(e);
        }
    });
})();
