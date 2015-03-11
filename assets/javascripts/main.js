//= require javascripts/jquery-2.1.3.min
//= require javascripts/bootstrap.min

$(function() {
  $('#select-themes label').on('click', function() {
    $('body').removeClass().addClass('theme-' + $('input', this).val());
  });

  $('#helper').on('click', function() {
    $('#help-content').toggle();
  });

  var createConfirm = function(e) {
    if(!confirm('are you ok?')) {
      e.stopPropagation();
      e.preventDefault();
    }
  }

  $('#create').on('click', createConfirm);
  $('#add-comment').on('click', createConfirm);
});
