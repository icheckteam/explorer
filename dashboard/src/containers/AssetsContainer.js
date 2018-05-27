import React, { Component } from 'react';
import { connect } from 'react-redux';
import { getAllAssets } from '../actions/asset'
import Assets from './components/Assets';
function mapStateToProps(state) {
  return {
    assets: state.assets.assets,
  };
}

class AssetsContainer extends Component {
  componentDidMount() {
    this.props.getAllAssets()
  }
  render() {
    return (
      <div>
        <Assets
          assets={this.props.assets}/>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  {getAllAssets}
)(AssetsContainer);