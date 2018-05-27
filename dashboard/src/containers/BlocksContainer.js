import React, { Component } from 'react';
import { connect } from 'react-redux';
import { getAllBlocks } from '../actions/block'
import Blocks from './components/Blocks'

function mapStateToProps(state) {
  return {
    blocks: state.blocks.blocks,
  };
}

class BlocksContainer extends Component {
  componentDidMount() {
    this.props.getAllBlocks()
  }
  render() {
    return (
      <div>
        <Blocks 
          blocks={this.props.blocks}/>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  { getAllBlocks }
)(BlocksContainer);