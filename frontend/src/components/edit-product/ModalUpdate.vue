<template>
     <!-- Modal HTML -->
  <div id="myModal" class="modal fade" tabindex="-1">
    <div class="modal-dialog">
      <div class="modal-content">
        <!-- header -->
        <div class="modal-header">
          <h5 class="modal-title">Function Group</h5>
          <button type="button" class="close" data-dismiss="modal">
            &times;
          </button>
        </div>

        <!-- body -->
        <div class="modal-body">
          <div class="modal-container">
            <input type="hidden" id="id" name="id" />

            <label for="name"><b>Name</b></label>
            <input
              type="text"
              placeholder="Enter Name"
              name="name"
              id="name"
              required
            />
            <br />
            <label for="name"><b>Total Member</b></label>
            <input
              type="text"
              placeholder="Enter Total Member"
              name="totalMember"
              id="totalMember"
              required
            />
            <br />
          </div>
        </div>

        <!-- footer -->
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">
            Cancel
          </button>
          <button type="button" class="btn btn-primary" onclick="save()">
            Save
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  
}
</script>