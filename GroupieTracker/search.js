$(document).ready(function() {
  $('#searchbar').on('submit', function(e) {
    e.preventDefault();
    var searchTerm = $('#name').val();
    $.ajax({
      type: 'POST',
      url: '/search',
      data: {searchTerm: searchTerm},
      success: function(response) {
        var results = $('#results');
        results.empty();
        $.each(response, function(i, item) {
          var li = $('<li>').text(item.name);
          results.append(li);
        });
      }
    });
  });
});