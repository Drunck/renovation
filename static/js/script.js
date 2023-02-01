$(window).on("load", function () {
    $("#preloader").fadeOut("slow", function () {
        $(this).remove();
    });
});

$(window).scroll(function () {
    if ($(window).scrollTop() > 60) {
        if ($(window).width() > 1200) {
            $("header").css({
                padding: "0.2rem 0",
                transition: "0.7s ease",
            });
        } else {
            $("header").css({
                padding: "0.5rem 0",
                transition: "0.7s ease",
            });
        }
    } else {
        $("header").css({
            padding: "0.8rem 0",
            transition: "0.7s ease",
        });
    }
});

$("a").on("click", function (e) {
    if (this.hash != "") {
        // e.preventDefault();
        const hash = this.hash;

        $("html, body").animate(
            {
                scrollTop: $(hash).offset().top,
            },
            800
        );
    }
});

$(".selectPackageButton").click(function () {
    let selectedValue = $(this).data("value");
    $("#validationCustom04").val(selectedValue);
});

// gallery modal

var Offcanvas = $("#offcanvas");
var bsOffcanvas = new bootstrap.Offcanvas(Offcanvas);

$("#offcanvas a").click(function () {
    bsOffcanvas.hide();
});
