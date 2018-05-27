import React, { Component } from 'react';
import { connect } from 'react-redux';

function mapStateToProps(state) {
  return {

  };
}

class TransactionsContainer extends Component {
  render() {
    return (
      <div>
        TransactionsContainer
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
)(TransactionsContainer);