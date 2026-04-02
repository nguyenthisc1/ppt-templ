document.addEventListener("DOMContentLoaded", function () {
  var toggle = document.querySelector("[data-nav-toggle]");
  var nav = document.querySelector("[data-nav]");
  var modal = document.querySelector("[data-contact-modal]");
  var openers = document.querySelectorAll("[data-open-contact]");
  var closers = document.querySelectorAll("[data-close-contact]");
  var slider = document.querySelector("[data-slider]");

  if (!toggle || !nav) {
    // Continue: the contact modal can still be available on pages without nav controls.
  } else {
    toggle.addEventListener("click", function () {
      var isOpen = nav.classList.toggle("nav--open");
      toggle.setAttribute("aria-expanded", String(isOpen));
    });
  }

  if (!modal) {
    return;
  }

  openers.forEach(function (opener) {
    opener.addEventListener("click", function () {
      if (typeof modal.showModal === "function") {
        modal.showModal();
      }
    });
  });

  closers.forEach(function (closer) {
    closer.addEventListener("click", function () {
      modal.close();
    });
  });

  modal.addEventListener("click", function (event) {
    var panel = modal.querySelector(".contact-modal__panel");
    if (!panel) {
      return;
    }

    if (!panel.contains(event.target)) {
      modal.close();
    }
  });

  if (!slider) {
    return;
  }

  var slides = Array.prototype.slice.call(slider.querySelectorAll("[data-slide]"));
  var dots = Array.prototype.slice.call(slider.querySelectorAll("[data-slider-dot]"));
  var prev = slider.querySelector("[data-slider-prev]");
  var next = slider.querySelector("[data-slider-next]");
  var currentIndex = 0;

  function renderSlider(index) {
    currentIndex = index;
    slides.forEach(function (slide, slideIndex) {
      slide.toggleAttribute("data-active", slideIndex === currentIndex);
    });
    dots.forEach(function (dot, dotIndex) {
      dot.toggleAttribute("data-active", dotIndex === currentIndex);
    });
  }

  if (prev) {
    prev.addEventListener("click", function () {
      renderSlider((currentIndex - 1 + slides.length) % slides.length);
    });
  }

  if (next) {
    next.addEventListener("click", function () {
      renderSlider((currentIndex + 1) % slides.length);
    });
  }

  dots.forEach(function (dot, dotIndex) {
    dot.addEventListener("click", function () {
      renderSlider(dotIndex);
    });
  });

  if (slides.length > 0) {
    renderSlider(0);
  }
});
