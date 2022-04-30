import React from 'react';
import PropTypes from 'prop-types';

require('./modal-list-selectable-item.scss');

function ModalListSelectableItem({
  name,
  selected,
  changeFunction,
  uuid,
}) {
  return (
    <div className="modal-list-selectable-item" data-uuid={uuid}>
      <input
        className="modal-checkbox"
        type="checkbox"
        checked={selected}
        onChange={changeFunction}
      />
      <span className="name-label" onClick={changeFunction}>{name}</span>
    </div>
  );
}

ModalListSelectableItem.propTypes = {
  uuid: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  selected: PropTypes.bool.isRequired,
  changeFunction: PropTypes.func.isRequired,
};

export default ModalListSelectableItem;
