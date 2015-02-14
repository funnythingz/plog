//= require javascripts/jquery-2.1.3.min
//= require javascripts/bootstrap.min

$(function() {
  $('body').addClass('theme-' + $('#select-themes input').val());
  $('#select-themes label').on('click', function() {
    $('body').removeClass().addClass('theme-' + $('input', this).val());
  });

  $('#helper').on('click', function() {
    $('#help-content').toggle();
  });
});
