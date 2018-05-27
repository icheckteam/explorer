import React, { Component } from 'react';
import { connect } from 'react-redux';
import { getAllAddresses } from '../actions/address'
import Addresses from './components/Addresses';
function mapStateToProps(state) {
  return {
    addresses: state.addresses.addresses,
  };
}

class AddressesContainer extends Component {
  componentDidMount() {
    this.props.getAllAddresses()
  }
  render() {
    return (
      <div>
        <Addresses
          addresses={this.props.addresses}
          />
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  { getAllAddresses }
)(AddressesContainer);