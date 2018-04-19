'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  isListItem: _propTypes2.default.bool,
  inset: _propTypes2.default.bool
};

var defaultProps = {
  isListItem: false
};

var ListDivider = function ListDivider(_ref) {
  var inset = _ref.inset,
      isListItem = _ref.isListItem,
      otherProps = _objectWithoutProperties(_ref, ['inset', 'isListItem']);

  var tag = isListItem ? 'li' : 'hr';
  var className = (0, _classnames2.default)('mdc-list-divider', {
    'mdc-list-divider--inset': inset
  }, otherProps.className);
  var props = isListItem ? { className: className, role: 'separator' } : { className: className };
  return (0, _react.createElement)(tag, props);
};
ListDivider.propTypes = propTypes;
ListDivider.defaultProps = defaultProps;
exports.default = ListDivider;