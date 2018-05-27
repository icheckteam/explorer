import React, { Component } from 'react';
import { connect } from 'react-redux';
import Transactions from './components/Transactions';
import { getAllTransactions } from '../actions/transaction'
function mapStateToProps(state) {
  return {
    transactions: state.transactions.transactions,
  };
}

class TransactionsContainer extends Component {
  componentDidMount() {
    this.props.getAllTransactions()
  }
  render() {
    return (
      <div>
        <Transactions
          transactions={this.props.transactions}
          />
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  { getAllTransactions }
)(TransactionsContainer);