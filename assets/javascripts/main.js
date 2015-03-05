//= require javascripts/jquery-2.1.3.min
//= require javascripts/bootstrap.min

$(function() {
  $('#select-themes label').on('click', function() {
    $('body').removeClass().addClass('theme-' + $('input', this).val());
  });

  $('#helper').on('click', function() {
    $('#help-content').toggle();
  });

  $('#create').on('click', function(e) {
    if(!confirm('are you ok?')) {
      e.stopPropagation();
      e.preventDefault();
    }
  });
});
