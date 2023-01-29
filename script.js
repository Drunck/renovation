$(window).scroll(function () {
    if ($(window).scrollTop() > 60) {
        $("header").css({
            padding: "0.2rem 0",
            transition: "0.7s ease",
        });
    } else {
        $("header").css({
            padding: "0.8rem 0",
            transition: "0.7s ease",
        });
    }
});

$("header a").on("click", function (e) {
    if (this.hash != "") {
        e.preventDefault();
        const hash = this.hash;

        $("html, body").animate(
            {
                scrollTop: $(hash).offset().top,
            },
            800
        );
    }
});
