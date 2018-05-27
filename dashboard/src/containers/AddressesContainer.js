import React, { Component } from 'react';
import { connect } from 'react-redux';

function mapStateToProps(state) {
  return {

  };
}

class AddressesContainer extends Component {
  render() {
    return (
      <div>
        AddressesContainer
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
)(AddressesContainer);